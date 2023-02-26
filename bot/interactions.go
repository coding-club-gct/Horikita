package bot

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func EventUserMenuSelectInteraction(ev discord.InteractionEvent) *api.InteractionResponse {
	_, ok := ev.Data.(*discord.SelectInteraction)
	fmt.Println(ok)
	resp := &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags:   api.EphemeralResponse,
			Content: option.NewNullableString("Event registered successfully"),
		},
	}
	if err := H.s.RespondInteraction(ev.ID, ev.Token, *resp); err != nil {
		fmt.Println("failed to send interaction callback:", err)
	}
	return resp
}
