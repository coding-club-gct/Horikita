package main

import (
	"log"
	"os"

	"github.com/joel-samuel-raj/Horikita/Bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if(err != nil) {
		log.Fatalf("Error ocurred while loading env file:  %s", err)
	}
	token := os.Getenv("token")
	bot.Connect(token)
}
