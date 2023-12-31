package v1

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "users-segments-service/docs/swagger"
	"users-segments-service/internal/usecase"
	"users-segments-service/pkg/logger"
)

// SetupRoutes -.
// Swagger spec:
// @title       Users segments service API
// @description Users segments
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func SetupRoutes(handler *fiber.App, u usecase.User, s usecase.Segment, us usecase.UsersSegment, ur usecase.Report, rd string, l logger.Interface) {

	handler.Get("/swagger/*", swagger.HandlerDefault)

	h := handler.Group("/v1")
	h.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	SetUserRoutes(h, u, l)
	SetSegmentRoutes(h, s, us, l)
	SetUsersSegmentsRoutes(h, u, s, us, ur, l)
	SetReportsRoutes(h, rd, l)
}
