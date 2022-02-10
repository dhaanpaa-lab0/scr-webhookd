package main

import (
	_ "github.com/dhaanpaa-lab0/scr-webhookd/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World"})
	})

	err := app.Listen(":3001")
	if err != nil {
		return
	}
}
