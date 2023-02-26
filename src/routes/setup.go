package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Route struct {
	Group   string
	Handler func(api fiber.Router)
}

var routes = []Route{
	{
		Group:   "/",
		Handler: HandleHelloWorld,
	},
	{
		Group:   "/events",
		Handler: HandleEvents,
	},
}

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	for _, route := range routes {
		route.Handler(api.Group(route.Group))
	}

}
