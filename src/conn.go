package src

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Conn () {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatalln("Couldnt get PORT")
	}
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":"+PORT)
}
