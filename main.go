package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joel-samuel-raj/Horikita/bot"
	"github.com/joel-samuel-raj/Horikita/database"
	route "github.com/joel-samuel-raj/Horikita/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	ENV := os.Getenv("ENV")
	if ENV == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error ocurred while loading env file:  %s", err)
		}
	}
	TOKEN := os.Getenv("TOKEN")
	PORT := os.Getenv("PORT")
	database.SqliteDBC()
	go bot.Conn(TOKEN)
	app := fiber.New(); route.SetupRoutes(app)
	app.Listen(":"+PORT)
}
