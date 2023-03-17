package main

import (
	"log"
	"os"

	"github.com/joel-samuel-raj/Horikita/bot"
	"github.com/joel-samuel-raj/Horikita/constants"
	"github.com/joho/godotenv"
)

func main() {

	constants.C.ServerURL = "http://localhost:1337"
	
	godotenv.Load()
	API_TOKEN := os.Getenv("API_TOKEN")
	if API_TOKEN == "" {
		log.Fatalln("No API_TOKEN given :(")
	}

	constants.C.ServerApiToken = API_TOKEN 
	bot.Conn()
}
