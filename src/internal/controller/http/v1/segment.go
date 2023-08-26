package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"users-segments-service/internal/controller/http/v1/middleware"
	"users-segments-service/internal/usecase"
	"users-segments-service/pkg/database"
	"users-segments-service/pkg/logger"
	"users-segments-service/pkg/validations"
)

type SegmentRoutes struct {
	s  usecase.Segment
	us usecase.UsersSegment
	l  logger.Interface
}

type createSegmentRequest struct {
	Slug string `json:"slug" binding:"required"  example:"AVITO_VOICE_MESSAGES"`
}

type createSegmentResponse struct {
	ID   uint   `json:"id" binding:"required"  example:"1"`
	Slug string `json:"slug" binding:"required"  example:"AVITO_VOICE_MESSAGES"`
}

type segmentsResponse struct {
	Segments []string `json:"segments" example:"AVITO_VOICE_MESSAGES"`
}

func SetSegmentRoutes(handler fiber.Router, s usecase.Segment, us usecase.UsersSegment, l logger.Interface) {
	r := &SegmentRoutes{
		s:  s,
		us: us,
		l:  l,
	}
	h := handler.Group("/segments")
	h.Post("", r.create)
	h.Get("", r.get)
	h.Delete("/:segmentSlug", middleware.GormTransaction(database.DB, l), r.delete)
}

// @Summary     Get
// @Description Returns all segments
// @ID          getSegments
// @Tags  	    segments
// @Accept      json
// @Produce     json
// @Success     200 {object} Response{data=segmentsResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /segments [get]
func (r *SegmentRoutes) get(c *fiber.Ctx) error {

	segments, err := r.s.GetSegments()
	if err != nil && err != gorm.ErrRecordNotFound {
		r.l.Error(err, "v1 - get - r.s.GetSegments")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get segments", nil, err)
	}

	response := make([]string, len(segments), len(segments))
	for i, segment := range segments {
		response[i] = segment.Slug
	}

	return OkResponse(c, fiber.StatusOK, "OK", &segmentsResponse{Segments: response})
}

// @Summary     Delete
// @Description Deletes segment and all users from it
// @ID          delete
// @Tags  	    segments
// @Accept      json
// @Produce     json
// @Param 		segmentSlug path string true "segment ID" example(1)
// @Success     200 {object} Response
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /segments/:segmentSlug [delete]
func (r *SegmentRoutes) delete(c *fiber.Ctx) error {

	segmentSlug := c.Params("segmentSlug")

	exists, err := r.s.SegmentExists(segmentSlug)
	if err != nil {
		r.l.Error(err, "v1 - delete - r.s.SegmentExists")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to check if segment exists", nil, err)
	}
	if !exists {
		return ErrorResponse(c, fiber.StatusBadRequest, fmt.Sprintf("Segment %s does not exist", segmentSlug), nil, ErrorSegmentDoesNotExist)
	}

	segment, err := r.us.SegmentBySlug(segmentSlug)
	if err != nil {
		r.l.Error(err, "v1 - delete - r.us.SegmentBySlug")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to get segment", nil, err)
	}

	err = r.us.DeleteAllUsersFromSegment(segment.ID)
	if err != nil {
		r.l.Error(err, "v1 - delete - r.us.DeleteAllUsersFromSegment")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete users from segment", nil, err)
	}

	err = r.s.Delete(segmentSlug)
	if err != nil {
		r.l.Error(err, "v1 - delete - r.s.Delete")
		return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete segment", nil, err)
	}

	return OkResponse(c, fiber.StatusOK, "Deleted", nil)
}

// @Summary     Create
// @Description Create a new segment
// @ID          create
// @Tags  	    segments
// @Accept      json
// @Produce     json
// @Param       request body createSegmentRequest true "Segment data"
// @Success     201 {object} Response{data=createSegmentResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /segments [post]
func (r *SegmentRoutes) create(c *fiber.Ctx) error {
	req := &createSegmentRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - create - c.BodyParser")
		return ErrorResponse(c, fiber.StatusBadRequest, MsgInvalidRequestBody, nil, err)
	}

	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return ErrorResponse(c, fiber.StatusBadRequest, MsgNotValid, errs, validations.ErrorValidationFailed)
	}

	res, err := r.s.Create(
		req.Slug,
	)

	if err != nil {
		if err == usecase.ErrorSegmentAlreadyExists {
			return ErrorResponse(c, fiber.StatusBadRequest, MsgSegmentAlreadyExists, nil, err)
		} else {
			r.l.Error(err, "v1 - create - r.s.Create")
			return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create segment", nil, err)
		}
	}

	responseData := &createSegmentResponse{
		ID:   res.ID,
		Slug: res.Slug,
	}

	return OkResponse(c, fiber.StatusCreated, "Segment created", responseData)
}
