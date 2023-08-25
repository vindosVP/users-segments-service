package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
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
	l  logger.Interface
}

type addUserToSegmentRequest struct {
	Segments []string `json:"segments" binding:"required"  example:"[AVITO_VOICE_MESSAGES, AVITO_PERFORMANCE_VAS]" validate:"required"`
}

type usersSegmentsResponse struct {
	UsersSegments []string `json:"usersSegments" example:"[AVITO_VOICE_MESSAGES, AVITO_PERFORMANCE_VAS]"`
}

func SetUsersSegmentsRoutes(handler fiber.Router, u usecase.User, s usecase.Segment, us usecase.UsersSegment, l logger.Interface) {
	r := &UsersSegmentsRoutes{
		u:  u,
		s:  s,
		us: us,
		l:  l,
	}
	h := handler.Group("/users")
	h.Post("/:userID/segments/add", middleware.GormTransaction(database.DB, l), r.add)
	h.Post("/:userID/segments/delete", middleware.GormTransaction(database.DB, l), r.delete)
}

// @Summary     Add to segments
// @Description Adds user to segments
// @ID          add
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body addUserToSegmentRequest true "123"
// @Success     200 {object} Response{data=usersSegmentsResponse, error=null}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users/:userID/segments/add [post]
func (usr *UsersSegmentsRoutes) add(c *fiber.Ctx) error {
	req := &addUserToSegmentRequest{}
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
				return ErrorResponse(c, fiber.StatusInternalServerError, fmt.Sprintf("User is already added to %s", slug), nil, err)
			}
			usr.l.Error(err, "v1 - add - usr.us.AddUserToSegment")
			return ErrorResponse(c, fiber.StatusInternalServerError, fmt.Sprintf("Failed to add user to %s", slug), nil, err)
		}
	}

	return OkResponse(c, fiber.StatusOK, "Added", nil)
}

func (usr *UsersSegmentsRoutes) delete(c *fiber.Ctx) error {
	return OkResponse(c, fiber.StatusOK, "Deleted", nil)
}
