package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
	"users-segments-service/internal/controller/http/v1/middleware"
	"users-segments-service/internal/usecase"
	"users-segments-service/pkg/database"
	"users-segments-service/pkg/logger"
	"users-segments-service/pkg/validations"
)

type UsersSegmentsRoutes struct {
	u  usecase.User
	s  usecase.Segment
	us usecase.UsersSegment
	ur usecase.Report
	l  logger.Interface
}

type modifyUsersSegmentsRequest struct {
	Segments []string `json:"segments" binding:"required"  example:"AVITO_VOICE_MESSAGES,AVITO_PERFORMANCE_VAS" validate:"required"`
}

type usersSegmentsResponse struct {
	UsersSegments []string `json:"usersSegments" example:"AVITO_VOICE_MESSAGES,AVITO_PERFORMANCE_VAS"`
}

type ReportResponse struct {
	FileLink string `json:"fileLink" example:"http://localhost:8080/v1/reports/80ef1ba7-1045-41aa-a8a2-4c0aba407baf"`
}

func SetUsersSegmentsRoutes(handler fiber.Router, u usecase.User, s usecase.Segment, us usecase.UsersSegment, ur usecase.Report, l logger.Interface) {
	r := &UsersSegmentsRoutes{
		u:  u,
		s:  s,
		us: us,
		ur: ur,
		l:  l,
	}
	h := handler.Group("/users")
	h.Get("/:userID/segments", r.get)
	h.Post("/:userID/segments/add", middleware.GormTransaction(database.DB, l), r.add)
	h.Post("/:userID/segments/delete", middleware.GormTransaction(database.DB, l), r.delete)
	h.Get("/:userID/segments/report", r.report)
}

// @Summary     Get users segments report
// @Description Returns link to a csv with user segments report
// @ID          getUsersSegmentsReport
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param 		userID path string true "user ID" example(1)
// @Param 		month query string true "month" example(8)
// @Param 		year query string true "year" example(2023)
// @Success     200 {object} Response{data=ReportResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/:userID/segments/report [get]
func (usr *UsersSegmentsRoutes) report(c *fiber.Ctx) error {

	month := c.Query("month")
	if month == "" {
		return ErrorResponse(c, fiber.StatusBadRequest, "No month specified", nil, nil)
	}
	year := c.Query("year")
	if year == "" {
		return ErrorResponse(c, fiber.StatusBadRequest, "No year specified", nil, nil)
	}

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse year", nil, nil)
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse month", nil, nil)
	}
	location, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to load location", nil, nil)
	}

	startTime := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 1, 0, location)
	endTime := startTime.AddDate(0, 1, 0)

	if !(monthInt > 0 && monthInt < 13) {
		return ErrorResponse(c, fiber.StatusBadRequest, "Wrong month", nil, nil)
	}

	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		usr.l.Error(err, "v1 - report - strconv.ParseUint")
		return ErrorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseID, nil, err)
	}

	userExists, err := usr.u.UserExistsByID(uint(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - report - usr.u.UserExistsByID")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to check if user exists", nil, err)
	}
	if !userExists {
		return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("User with ID %d does not exist", userID), nil, ErrorUserDoesNotExist)
	}

	report, err := usr.us.Report(uint(userID), startTime, endTime)
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - report - usr.us.Report")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get report", nil, err)
	}

	reportFileName, err := usr.ur.SaveReport(report)
	if err != nil {
		usr.l.Error(err, "v1 - report - usr.ur.SaveReport")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to save report", nil, err)
	}

	resp := &ReportResponse{
		FileLink: fmt.Sprintf("http://localhost:8080/v1/reports/%s", reportFileName),
	}

	return OkResponse(c, fiber.StatusOK, "OK", resp)
}

// @Summary     Get users segments
// @Description Returns all users segments
// @ID          getUsersSegments
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param 		userID path string true "user ID" example(1)
// @Success     200 {object} Response{data=usersSegmentsResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/:userID/segments [get]
func (usr *UsersSegmentsRoutes) get(c *fiber.Ctx) error {

	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		usr.l.Error(err, "v1 - get - strconv.ParseUint")
		return ErrorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseID, nil, err)
	}

	userExists, err := usr.u.UserExistsByID(uint(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - get - usr.u.UserExistsByID")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to check if user exists", nil, err)
	}
	if !userExists {
		return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("User with ID %d does not exist", userID), nil, ErrorUserDoesNotExist)
	}

	segments, err := usr.us.GetUsersSegments(uint(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - get - usr.us.GetUsersSegments")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get users segments", nil, err)
	}

	response := &usersSegmentsResponse{
		UsersSegments: segments,
	}

	return OkResponse(c, fiber.StatusOK, "OK", response)
}

// @Summary     Add to segments
// @Description Adds user to segments
// @ID          add
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body modifyUsersSegmentsRequest true "Segments to add"
// @Param 		userID path string true "user ID" example(1)
// @Success     200 {object} Response{data=usersSegmentsResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/:userID/segments/add [post]
func (usr *UsersSegmentsRoutes) add(c *fiber.Ctx) error {
	req := &modifyUsersSegmentsRequest{}
	if err := c.BodyParser(req); err != nil {
		usr.l.Error(err, "v1 - add - c.BodyParser")
		return ErrorResponse(c, fiber.StatusBadRequest, MsgInvalidRequestBody, nil, err)
	}

	isValid, errs := validations.UniversalValidation(req)
	if !isValid {
		return ErrorResponse(c, fiber.StatusBadRequest, MsgNotValid, errs, validations.ErrorValidationFailed)
	}

	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		usr.l.Error(err, "v1 - add - strconv.ParseUint")
		return ErrorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseID, nil, err)
	}

	userExists, err := usr.u.UserExistsByID(uint(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - add - usr.u.UserExistsByID")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to check if user exists", nil, err)
	}
	if !userExists {
		return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("User with ID %d does not exist", userID), nil, ErrorUserDoesNotExist)
	}

	for _, slug := range req.Segments {
		segmentExists, err := usr.s.SegmentExists(slug)
		if err != nil && err != gorm.ErrRecordNotFound {
			usr.l.Error(err, "v1 - add - usr.s.SegmentExists")
			return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to check if segment exists", nil, err)
		}
		if !segmentExists {
			return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("Segment with slug %s does not exist", slug), nil, ErrorSegmentDoesNotExist)
		}
	}

	for _, slug := range req.Segments {
		_, err := usr.us.AddUserToSegment(uint(userID), slug)
		if err != nil {
			if err == usecase.ErrorUserAlreadyAdded {
				return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("User is already added to %s", slug), nil, err)
			}
			usr.l.Error(err, "v1 - add - usr.us.AddUserToSegment")
			return ErrorResponse(c, fiber.StatusInternalServerError, fmt.Sprintf("Failed to add user to %s", slug), nil, err)
		}
	}

	segments, err := usr.us.GetUsersSegments(uint(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - add - usr.us.GetUsersSegments")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get users segments", nil, err)
	}

	response := &usersSegmentsResponse{
		UsersSegments: segments,
	}

	return OkResponse(c, fiber.StatusOK, "Added", response)
}

// @Summary     Delete from segment
// @Description Deletes user from segment
// @ID          deleteFromSegment
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body modifyUsersSegmentsRequest true "Segments to delete"
// @Param 		userID path string true "user ID" example(1)
// @Success     200 {object} Response{data=usersSegmentsResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/:userID/segments/delete [post]
func (usr *UsersSegmentsRoutes) delete(c *fiber.Ctx) error {

	req := &modifyUsersSegmentsRequest{}
	if err := c.BodyParser(req); err != nil {
		usr.l.Error(err, "v1 - delete - c.BodyParser")
		return ErrorResponse(c, fiber.StatusBadRequest, MsgInvalidRequestBody, nil, err)
	}

	isValid, errs := validations.UniversalValidation(req)
	if !isValid {
		return ErrorResponse(c, fiber.StatusBadRequest, MsgNotValid, errs, validations.ErrorValidationFailed)
	}

	userID, err := strconv.ParseUint(c.Params("userID"), 10, 32)
	if err != nil {
		usr.l.Error(err, "v1 - delete - strconv.ParseUint")
		return ErrorResponse(c, fiber.StatusInternalServerError, MsgFailedToParseID, nil, err)
	}

	userExists, err := usr.u.UserExistsByID(uint(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - delete - usr.u.UserExistsByID")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to check if user exists", nil, err)
	}
	if !userExists {
		return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("User with ID %d does not exist", userID), nil, ErrorUserDoesNotExist)
	}

	for _, slug := range req.Segments {
		segmentExists, err := usr.s.SegmentExists(slug)
		if err != nil && err != gorm.ErrRecordNotFound {
			usr.l.Error(err, "v1 - delete - usr.s.SegmentExists")
			return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to check if segment exists", nil, err)
		}
		if !segmentExists {
			return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("Segment with slug %s does not exist", slug), nil, ErrorSegmentDoesNotExist)
		}
	}

	for _, slug := range req.Segments {
		err := usr.us.DeleteUsersSegment(uint(userID), slug)
		if err != nil {
			if err == usecase.ErrorUserIsNotAddedToSegment {
				return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("User is not added to segment %s", slug), nil, err)
			}
			usr.l.Error(err, "v1 - delete - usr.us.AddUserToSegment")
			return ErrorResponse(c, fiber.StatusInternalServerError, fmt.Sprintf("Failed to delete user from segment %s", slug), nil, err)
		}
	}

	segments, err := usr.us.GetUsersSegments(uint(userID))
	if err != nil && err != gorm.ErrRecordNotFound {
		usr.l.Error(err, "v1 - add - usr.us.GetUsersSegments")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get users segments", nil, err)
	}

	response := &usersSegmentsResponse{
		UsersSegments: segments,
	}

	return OkResponse(c, fiber.StatusOK, "Deleted", response)
}
