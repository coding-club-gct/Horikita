package route

import "github.com/gofiber/fiber/v2"

func HandleHelloWorld(api fiber.Router) {
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World from Horikita fiber routine")
	})
}
