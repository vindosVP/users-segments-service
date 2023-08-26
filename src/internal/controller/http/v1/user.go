package v1

import (
	"github.com/gofiber/fiber/v2"
	"users-segments-service/internal/usecase"
	"users-segments-service/pkg/logger"
	"users-segments-service/pkg/validations"
)

type UserRoutes struct {
	u usecase.User
	l logger.Interface
}

type registerUserRequest struct {
	Email    string `json:"email" binding:"required"  example:"vadiminmail@gmail.com" validate:"required,email"`
	Name     string `json:"name" binding:"required"  example:"Vadim" validate:"required"`
	LastName string `json:"lastName" binding:"required"  example:"Valov" validate:"required"`
}

type registerUserResponse struct {
	ID       uint   `json:"id" binding:"required"  example:"1"`
	Email    string `json:"email" binding:"required"  example:"vadiminmail@gmail.com"`
	Name     string `json:"name" binding:"required"  example:"Vadim"`
	LastName string `json:"lastName" binding:"required"  example:"Valov"`
}

func SetUserRoutes(handler fiber.Router, u usecase.User, l logger.Interface) {
	r := &UserRoutes{
		u: u,
		l: l,
	}
	h := handler.Group("/users")
	h.Post("", r.register)
}

// @Summary     Register
// @Description Register a new user
// @ID          register
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Param       request body registerUserRequest true "User data"
// @Success     201 {object} Response{data=registerUserResponse}
// @Failure     400 {object} Response
// @Failure     500 {object} Response
// @Router      /users [post]
func (r *UserRoutes) register(c *fiber.Ctx) error {
	req := &registerUserRequest{}
	if err := c.BodyParser(req); err != nil {
		r.l.Error(err, "v1 - register - c.BodyParser")
		return ErrorResponse(c, fiber.StatusBadRequest, MsgInvalidRequestBody, nil, err)
	}

	isValid, errs := validations.UniversalValidation(req)

	if !isValid {
		return ErrorResponse(c, fiber.StatusBadRequest, MsgNotValid, errs, validations.ErrorValidationFailed)
	}

	res, err := r.u.Register(
		req.Email,
		req.Name,
		req.LastName,
	)

	if err != nil {
		if err == usecase.ErrorUserAlreadyExists {
			return ErrorResponse(c, fiber.StatusBadRequest, MsgUserAlreadyExists, nil, err)
		} else {
			r.l.Error(err, "http - v1 - r.u.Register")
			return ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user", nil, err)
		}
	}

	responseData := &registerUserResponse{
		ID:       res.ID,
		Email:    res.Email,
		Name:     res.Name,
		LastName: res.LastName,
	}

	return OkResponse(c, fiber.StatusCreated, "User created", responseData)
}
