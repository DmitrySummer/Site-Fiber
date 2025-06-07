package pages

import (
	tadaprot "dl/new-web-site/pkg/tadaptop"
	"dl/new-web-site/views"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router fiber.Router
	log    *slog.Logger
}

func NewHandler(router fiber.Router, log *slog.Logger) *PagesHandler {
	h := &PagesHandler{
		router: router,
		log:    log,
	}

	h.router.Get("/", h.pages)
	return h
}

func (h *PagesHandler) pages(c *fiber.Ctx) error {
	components := views.Main()
	return tadaprot.Render(c, components)
}


