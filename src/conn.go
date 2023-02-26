package src

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	route "github.com/joel-samuel-raj/Horikita/src/routes"
)

func Conn() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatalln("Couldnt get PORT")
	}
	app := fiber.New()
	route.SetupRoutes(app)
	app.Listen(":" + PORT)
}
