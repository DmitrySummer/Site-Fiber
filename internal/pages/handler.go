package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router fiber.Router
	log    *slog.Logger
}

type Category struct {
	Name string
}

var categories = []Category{
	{Name: "Еда"},
	{Name: "Животные"},
	{Name: "Машины"},
	{Name: "Спорт"},
	{Name: "Музыка"},
	{Name: "Технологии"},
}

func NewHandler(router fiber.Router, log *slog.Logger) *PagesHandler {
	h := &PagesHandler{
		router: router,
		log:    log,
	}

	api := h.router.Group("/pages")
	api.Get("/", h.pages)

	return h
}

func (h *PagesHandler) pages(c *fiber.Ctx) error {
	return c.Render("pages", categories)
}
