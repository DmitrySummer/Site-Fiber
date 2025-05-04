package main

import (
	"dl/new-web-site/config"
	"dl/new-web-site/internal/pages"
	"dl/new-web-site/pkg/logger"

	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Init()

	// Создаём конфиг и логгер
	logConfig := config.NewLogConfig()
	log := logger.NewLogger(logConfig)

	app := fiber.New()

	// Конфигурация middleware для логирования
	app.Use(slogfiber.New(log))

	pages.NewHandler(app)

	if err := app.Listen(":3000"); err != nil {
		log.Error("failed to start server", "error", err)
	}
}
