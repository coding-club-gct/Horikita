package bot

import (
	"context"
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joel-samuel-raj/Horikita/Events"
)

func Connect(token string) *state.State {

	log.Println("Connect")

	state := state.New(token)

	state.AddIntents(gateway.IntentGuilds)
	err := state.Connect(context.Background())
	if err != nil {
		log.Println("Error while connecting to the gateway ", err)
	}
	events.Tester()
	return state
}
