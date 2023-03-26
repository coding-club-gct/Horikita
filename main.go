package main

import (
	"log"
	"os"

	"github.com/joel-samuel-raj/Horikita/bot"
	"github.com/joel-samuel-raj/Horikita/constants"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	constants.C.Strings = map[string]string{}
	for _, token := range []string{"API_TOKEN", "BOT_TOKEN", "SERVER_URL"} {
		TOKEN := os.Getenv(token)
		if TOKEN == "" {
			log.Fatalln("No " + token + " was given")
			return 
			} else {
				constants.C.Strings[token] = TOKEN
			}
		}
	bot.Conn()
}
