package core

import (
	"fmt"
	"os"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

func Bootstrap() *fiber.App {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	
	logger.Info().
		Str("name", "startup").
		Str("version", fiber.Version).
		Str("url", "http://127.0.0.1:3000").
		Str("pid", fmt.Sprintf("%d", os.Getpid())).
		Msg("server started")
		
	app := fiber.New(fiber.Config {
		DisableStartupMessage: true,
	})

	httpLogger := logger.With().Str("name", "http").Logger()

	app.Use(requestid.New(requestid.Config{
		Header: "x-request-id",
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &httpLogger,
		Fields: []string{
			fiberzerolog.FieldRequestID,
			fiberzerolog.FieldStatus,
			fiberzerolog.FieldLatency,
			fiberzerolog.FieldMethod,
			fiberzerolog.FieldPath,
		},
	}))

	app.Get("/", HandleHello)

	return app
}
