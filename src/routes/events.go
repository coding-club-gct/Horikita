package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joel-samuel-raj/Horikita/bot"
	types "github.com/joel-samuel-raj/Horikita/types"
)

func HandleEvents(api fiber.Router) {
	api.Post("/new", func(c *fiber.Ctx) error {
		event := new(types.Event)
		if err := c.BodyParser(&event); err != nil {
			return err
		}
		bot.PostEvent(event)
		return nil
	})
}
