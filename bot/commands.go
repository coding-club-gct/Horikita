package bot

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
)

var Commands = []api.CreateCommandData{
	{
		Name: "ping",
		Description: "ping pong!",
	},
	{
		Name: "thonk",
		Description: "thonks",
	},
	{
		Name: "echo",
		Description: "echos option",
		Options: []discord.CommandOption{
			&discord.StringOption{
				OptionName: "argument",
				Description: "whats echoed back",
				Required: true,
			},
		},
	},
}
