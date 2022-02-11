package main

import (
	"github.com/dhaanpaa-lab0/scr-webhookd/config"
	"github.com/dhaanpaa-lab0/scr-webhookd/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Scripted WebHookD",
		ServerHeader: config.GetServerHeader(),
	})

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World", "system_root": 123})
	})

	app.Post("/s/:name", func(c *fiber.Ctx) error {
		scriptKey := c.Params("name")
		tempFileName := config.GetTempFile()
		utils.WriteStrToFile(tempFileName, string(c.Body()))
		return c.JSON(fiber.Map{"scriptName": scriptKey, "message": config.ExecScript(scriptKey, tempFileName)})
	})

	log.Println("Webhook config root ......... : " + config.GetSystemRoot())
	log.Println("Webhook scripts root ........ : " + config.GetSystemRootScriptsPath())
	err := app.Listen(config.GetListenAddress())
	if err != nil {
		return
	}

}
