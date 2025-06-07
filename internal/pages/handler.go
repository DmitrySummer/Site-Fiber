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
	h.router.Get("/register", h.registerGet)
	h.router.Post("/register", h.registerPost)
	return h
}

func (h *PagesHandler) pages(c *fiber.Ctx) error {
	components := views.Main()
	return tadaprot.Render(c, components)
}

func (h *PagesHandler) registerGet(c *fiber.Ctx) error {
	h.log.Info("Открыта страница /register")
	return tadaprot.Render(c, views.Register())
}

func (h *PagesHandler) registerPost(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	h.log.Info("Register form submitted",
		"name", name,
		"email", email,
		"password", password,
	)

	return c.SendString("Регистрация прошла успешно!")
}
