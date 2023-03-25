package bot

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
)

var Commands = []api.CreateCommandData{
	{
		Name: "team-create",
		Description: "Team Leaders can use this slash command to create a team for an Event",
		Options: []discord.CommandOption{
			&discord.StringOption{
				OptionName: "name",
				Description: "Enter Team Name",
				Required: true,
				MaxLength: option.NewInt(20),
			},
		},
	},
}
