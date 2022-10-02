package commands

import (
	"log"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

func Ping (session *session.Session) {
	session.AddHandler(func (c *gateway.MessageCreateEvent) {
		log.Println(c.Author.Username, "has sent", c.Content)
	})
}