package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"users-segments-service/config"
	"users-segments-service/internal/clients"
	v1 "users-segments-service/internal/controller/http/v1"
	"users-segments-service/internal/usecase"
	"users-segments-service/internal/usecase/segment_repo"
	"users-segments-service/internal/usecase/user_repo"
	"users-segments-service/internal/usecase/usersSegment_repo"
	"users-segments-service/pkg/database"
	"users-segments-service/pkg/logger"
	"users-segments-service/pkg/validations"
)

func Run(config *config.Config) {

	lg := logger.New(config.Log.Level)

	db, err := database.NewGorm(config.DB)
	if err != nil {
		lg.Fatal(fmt.Errorf("app - Run - database.NewGorm: %w", err))
	}
	if err := validations.InitValidations(); err != nil {
		lg.Fatal(fmt.Errorf("app - Run - validations.InitValidations: %w", err))
	}

	handler := fiber.New()

	handler.Use(fiberlog.New(fiberlog.Config{
		TimeZone: "Europe/Moscow",
		Format:   "[${time}] ${locals:request-id} ${status} - ${latency} ${method} ${path}​\n",
	}))

	handler.Use(requestid.New(requestid.Config{
		Header:     "X-Request-ID",
		ContextKey: "request-id",
	}))

	userUseCase := usecase.NewUserUseCase(user_repo.New(db))
	segmentUseCase := usecase.NewSegmentUseCase(segment_repo.New(db))
	usersSegmentUseCase := usecase.NewUsersSegmentUseCase(usersSegment_repo.New(db))
	pastebinClient, err := clients.NewPastebinClient(config.App.PastebinLogin, config.App.PastebinPwd, config.App.PastebinToken)
	if err != nil {
		lg.Fatal("Failed to setup pastebin client")
	}

	v1.SetupRoutes(handler, userUseCase, segmentUseCase, usersSegmentUseCase, pastebinClient, lg)
	if err := handler.Listen(":" + config.App.Port); err != nil {
		lg.Fatal(fmt.Errorf("app - Run - handler.Listen: %w", err))
	}
}
