package main

import (
	"log"
	"os"

	"github.com/joel-samuel-raj/Horikita/bot"
	"github.com/joel-samuel-raj/Horikita/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error ocurred while loading env file:  %s", err)
	}
	token := os.Getenv("token")
	database.SqliteDBC()
	bot.Connect("Bot " + token)
}
