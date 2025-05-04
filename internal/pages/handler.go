package pages

import (
	"dl/new-web-site/config"
	"dl/new-web-site/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type PagesHandler struct {
	router fiber.Router
	log    *slog.Logger
}

func NewHandler(router fiber.Router) {
	// Создаём конфиг и логгер
	logConfig := config.NewLogConfig()
	log := logger.NewLogger(logConfig)

	h := &PagesHandler{
		router: router,
		log:    log,
	}

	api := h.router.Group("/pages")

	// Добавляем middleware для логирования
	api.Use(func(c *fiber.Ctx) error {
		h.log.Info("incoming request",
			"method", c.Method(),
			"path", c.Path(),
			"ip", c.IP(),
		)
		return c.Next()
	})

	api.Get("/", h.pages)

}

func (h *PagesHandler) pages(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Limit params didn't undefined")
}
