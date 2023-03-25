package bot

import (
	"context"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joho/godotenv"
)

type Handler struct {
	s *state.State
}

var H Handler

func Conn() {
	godotenv.Load()
	TOKEN := os.Getenv("TOKEN")
	if TOKEN == "" {
		log.Fatalln("No TOKEN given :(")
	}
	H.s = state.New("Bot " + TOKEN)
	H.s.AddInteractionHandler(&H)
	// add remaining handlers
	for _, handler := range Handlers {
		H.s.AddHandler(handler)
	}
	H.s.AddIntents(gateway.IntentGuildMembers)
	H.s.AddIntents(gateway.IntentGuilds)
	H.s.AddIntents(gateway.IntentGuildMessages)
	H.s.AddIntents(gateway.IntentDirectMessages)
	H.s.AddHandler(func(a *gateway.ReadyEvent) {
		me, _ := H.s.Me()
		log.Println("connected to the gateway as", me.Tag())
	})

	if err := overwriteCommands(H.s); err != nil {
		log.Fatalln("cannot update commands: ", err)
	}

	if err := H.s.Connect(context.Background()); err != nil {
		log.Fatalln("cannot connect:", err)
	}
	defer H.s.Close()
}

