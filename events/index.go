package events

import (
	"fmt"
	"strings"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

type eventsDictionaryType struct {
	name    string
	intent  gateway.Intents
	handler interface{}
	// commands [] interface {}
}

var emptyIntent gateway.Intents

var eventsDictionary = [...]eventsDictionaryType{
	{
		name:    "Message Create",
		intent:  gateway.IntentGuildMessages,
		handler: func() {},
		// commands: ,
	},
}

func InitEvents(state *state.State) {
	state.AddIntents(gateway.IntentGuildMessages)
	state.AddIntents(gateway.IntentGuilds)
	state.AddHandler(func(message *gateway.MessageCreateEvent) {
		if message.Author.Bot {
			return
		}
		prompts := [3]string{"i am", "i'm", "i'am"}
		for _, word := range prompts {
			lCont := strings.ToLower(message.Content)
			tmp := strings.SplitN(lCont, word, 2)
			if len(tmp) == 2 {
				cId := message.ChannelID
				_, err := state.SendMessage(cId, "Hello there, "+strings.TrimSpace(tmp[1])+" ! This is Horikita 😃")
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		}
	})
}


