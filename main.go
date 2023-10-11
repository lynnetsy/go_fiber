package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		log.Info("Route Hello World")

		return c.SendString("Hello World")
	})

	app.Listen(":8010")

}
