package main

import (
	"dl/new-web-site/config"
	"dl/new-web-site/internal/pages"
	"dl/new-web-site/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/slogfiber"
)

func main() {
	config.Init()

	// Создаём конфиг и логгер
	logConfig := config.NewLogConfig()
	log := logger.NewLogger(logConfig)

	app := fiber.New()

	// Конфигурация middleware для логирования
	app.Use(slogfiber.New(log))

	// Дополнительное логирование с помощью slog
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		log.Info("request completed",
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"duration", duration,
			"ip", c.IP(),
		)
		return err
	})

	pages.NewHandler(app)

	if err := app.Listen(":3000"); err != nil {
		log.Error("failed to start server", "error", err)
	}
}
