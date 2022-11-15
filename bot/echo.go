package bot

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

func (H *Handler) Echo (ev *discord.InteractionEvent, data *discord.CommandInteraction) *api.InteractionResponse {
	var options struct {
		Arg string `discord:"argument"`
	}
	if err := data.Options.Unmarshal(&options); err != nil {
		return ErrorResponse(err)
	}
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Content: option.NewNullableString(options.Arg),		
		},
	}
}
