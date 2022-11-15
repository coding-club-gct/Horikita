package bot

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (H *Handler) Thonk ( ev *discord.InteractionEvent, _ *discord.CommandInteraction ) *api.InteractionResponse {
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Content: option.NewNullableString("https://tenor.com/view/thonk-thinking-sun-thonk-sun-thinking-sun-gif-14999983"),
		},
	}
}
