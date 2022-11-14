package bot

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Handler struct {
	s *state.State	
}

var H Handler
func Conn () {
	godotenv.Load()
	TOKEN := os.Getenv("TOKEN")
	if TOKEN == "" {
		log.Fatalln("No TOKEN given :(")
	}
	H.s = state.New("Bot " + TOKEN)
	H.s.AddInteractionHandler(&H)
	H.s.AddIntents(gateway.IntentGuilds)
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

func overwriteCommands(s *state.State) error {
	app, err := s.CurrentApplication()
	if err != nil {
		return errors.Wrap(err, "cannot get current app ID")
	}

	_, err = s.BulkOverwriteCommands(app.ID, Commands)
	return err
}

func errorResponse(err error) *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Content: option.NewNullableString("**Error:** " + err.Error()),
			Flags: discord.EphemeralMessage,
			AllowedMentions: &api.AllowedMentions{ /* none */ },
		},
	}
}

func deferResponse(flags discord.MessageFlags) *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.DeferredMessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags: flags,
		},
	}
}

func (H *Handler) HandleInteraction (ev *discord.InteractionEvent) *api.InteractionResponse {
	switch data := ev.Data.(type) {
		case *discord.CommandInteraction:
			switch data.Name {
				case "ping":
					return H.Ping(ev, data)
				default:
					return errorResponse(fmt.Errorf("unknown command %q", data.Name))
			}
		default:
			return errorResponse(fmt.Errorf("unknown interaction %T", ev.Data))
	}
}







