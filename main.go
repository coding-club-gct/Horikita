package main

import (
	"github.com/joel-samuel-raj/Horikita/bot"
	"github.com/joel-samuel-raj/Horikita/src"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	go src.Conn()
	bot.Conn()
}
