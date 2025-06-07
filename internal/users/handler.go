package users

import (
	tadaprot "dl/new-web-site/pkg/tadaptop"
	"dl/new-web-site/pkg/validator"
	"dl/new-web-site/views/components"
	"log/slog"
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	registerGroup := h.router.Group("/api/register")
	registerGroup.Post("/", h.createUsers)
	return h
}

func (h *UsersHandler) createUsers(c *fiber.Ctx) error {
	form := UsersCreateForm{
		Email:    c.FormValue("email"),
		Name:     c.FormValue("name"),
		Password: c.FormValue("password"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Почта не задана или введена не верно"},
	)
	time.Sleep(time.Second * 1)
	if len(errors.Errors) > 0 {
		component := components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadaprot.Render(c, component)
	}
	component := components.Notification("Регистрация прошла успешно!", components.NotificationSuccess)

	h.log.Info("Register form submitted",
		"name", form.Name,
		"email", form.Email,
		"password", form.Password,
	)

	return tadaprot.Render(c, component)
}
