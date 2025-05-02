package main

import (
	"dl/new-web-site/internal/pages"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	pages.NewHandler(app)

	app.Listen(":3000")

}
