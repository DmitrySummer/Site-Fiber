package main

import (
	"dl/new-web-site/config"
	"dl/new-web-site/internal/pages"
	"dl/new-web-site/internal/users"
	"dl/new-web-site/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()

	logConfig := config.NewLogConfig()
	log := logger.NewLogger(logConfig)

	app := fiber.New()

	app.Static("/public", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 24 * 60 * 60 * time.Second,
	})

	app.Use(slogfiber.New(log))

	app.Use(func(c *fiber.Ctx) error {
		log.Info("incoming request",
			"method", c.Method(),
			"path", c.Path(),
			"ip", c.IP(),
		)
		return c.Next()
	})

	pages.NewHandler(app, log)
	users.NewHandler(app, log )

	if err := app.Listen(":3000"); err != nil {
		log.Error("failed to start server", "error", err)
	}
}

