package pages

import "github.com/gofiber/fiber/v2"

type PagesHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := &PagesHandler{
		router: router,
	}
	api := h.router.Group("/pages")
	api.Get("/", h.pages)

}

func (h *PagesHandler) pages(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "Limit params didn't undefined")
}
