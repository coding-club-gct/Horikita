package bot

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
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
			},
		},
	},
	{
		Name: "team-join",
		Description: "Team Members can use this slash command to join a team for an Event",
	},
}
