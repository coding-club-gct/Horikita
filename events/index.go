package events

import (
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

type dictionaryType struct {
	Intent gateway.Intents
	handler interface {}
}

var eventDictionary = [...] dictionaryType {
	{
		Intent: gateway.IntentGuildMessages,
		handler: func (c *gateway.MessageCreateEvent) {
			log.Println(c.Author)
		},
	},
}

func InitEvents (session *session.Session) {
	for _, event := range eventDictionary {
		session.AddIntents(event.Intent)
		session.AddHandler(event.handler)
	}
}