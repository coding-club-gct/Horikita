package main

import (
	"github.com/joel-samuel-raj/Horikita/bot"
	"github.com/joel-samuel-raj/Horikita/database"
	"github.com/joel-samuel-raj/Horikita/src"
)

func main () {
	database.Conn()
	go src.Conn()
	bot.Conn()
}










