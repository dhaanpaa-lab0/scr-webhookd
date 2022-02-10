package main

import (
	"github.com/dhaanpaa-lab0/scr-webhookd/config"
	_ "github.com/dhaanpaa-lab0/scr-webhookd/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Scripted WebHookD",
		ServerHeader: config.GetServerHeader(),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World", "system_root": 123})
	})

	log.Println("Webhook config root: " + config.GetSystemRoot())
	err := app.Listen(config.GetListenAddress())
	if err != nil {
		return
	}
}
