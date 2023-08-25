package v1

import (
	"github.com/gofiber/fiber/v2"
	"users-segments-service/internal/usecase"
	"users-segments-service/pkg/logger"
	"users-segments-service/pkg/validations"
)

type SegmentRoutes struct {
	s usecase.Segment
	l logger.Interface
}

type createSegmentRequest struct {
	Slug string `json:"slug" binding:"required"  example:"AVITO_VOICE_MESSAGES"`
}

type createSegmentResponse struct {
	ID   uint   `json:"id" binding:"required"  example:"1"`
	Slug string `json:"slug" binding:"required"  example:"AVITO_VOICE_MESSAGES"`
}

func SetSegmentRoutes(handler fiber.Router, s usecase.Segment, l logger.Interface) {
	r := &SegmentRoutes{
		s: s,
		l: l,
	}
	h := handler.Group("/segments")
	h.Post("", r.create)
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
