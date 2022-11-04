package bot

import (
	"context"
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

type botInstance struct {
	bot *session.Session
}

var Bot botInstance

func Conn (token string) {
	s := session.New("Bot " + token)
	s.AddIntents(gateway.IntentDirectMessages)
	s.AddIntents(gateway.IntentGuildMessages)
	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("Failed to connect: ", err)
	}
	defer s.Close()
	u, _ := s.Me()
	log.Println("Started as", u.Username)
	Bot = botInstance{bot: s}
}
