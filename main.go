package main

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func main() {
	app := fiber.New()

	app.Get("/:text?", func(c *fiber.Ctx) error {
		if c.Params("text") != "" {
			return c.SendString(strings.ToUpper(c.Params("text")))
		}
		return c.SendString("Demo App")
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
