package main

import (
	_ "github.com/dhaanpaa-lab0/scr-webhookd/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World", "system_root": viper.GetString("system_root")})
	})

	err := app.Listen(viper.GetString("listen_address"))
	if err != nil {
		return
	}
}
