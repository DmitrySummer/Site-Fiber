package users

import (
	tadaprot "dl/new-web-site/pkg/tadaptop"
	"dl/new-web-site/views"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	router fiber.Router
	log    *slog.Logger
}

func NewHandler(router fiber.Router, log *slog.Logger) *UsersHandler {
	h := &UsersHandler{
		router: router,
		log:    log,
	}
	registerGroup := h.router.Group("/register")
	registerGroup.Get("/", h.getUsers)
	registerGroup.Post("/", h.createUsers)
	return h
}

func (h *UsersHandler) getUsers(c *fiber.Ctx) error {
	h.log.Info("Открыта страница /register")
	return tadaprot.Render(c, views.Register())
}

func (h *UsersHandler) createUsers(c *fiber.Ctx) error {
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
