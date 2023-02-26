package bot

import (
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joel-samuel-raj/Horikita/types"
)

func PostEvent(event *types.Event) {
	_, err := H.s.Client.SendMessageComplex(893577313483096151, api.SendMessageData{
		Content: event.Name + "\n" + event.Description,
		Embeds: []discord.Embed{
			{
				Image: &discord.EmbedImage{
					URL: "https://imgs.search.brave.com/W9MyE0Q5eutjuWtJFJZoUIIXHAeldVSvABr0gCbcFtk/rs:fit:720:720:1/g:ce/aHR0cHM6Ly9lbS53/YXR0cGFkLmNvbS83/YTlmNGI1NzQxNTc4/MjU2MDhjMGZjM2Vl/ZWY4OGNjNTA5ZWU1/ZjE4LzY4NzQ3NDcw/NzMzYTJmMmY3MzMz/MmU2MTZkNjE3YTZm/NmU2MTc3NzMyZTYz/NmY2ZDJmNzc2MTc0/NzQ3MDYxNjQyZDZk/NjU2NDY5NjEyZDcz/NjU3Mjc2Njk2MzY1/MmY1Mzc0NmY3Mjc5/NDk2ZDYxNjc2NTJm/NGQ1YTQzNDM0ZDY0/NGI2NDQxNmQzMzc5/NmM2NzNkM2QyZDMz/MzMzMTJlMzEzNTYx/MzY2NjMzMzUzNTM2/MzQzMzM0NjEzMDM3/MzMzOTM5MzYzNTMx/MzEzMTM5MzEzNDM1/MzMyZTZhNzA2Nw",
				},
			},
		},
		Components: *discord.ComponentsPtr(
			&discord.UserSelectComponent{
				CustomID:    "EventUserSelectMenu",
				ValueLimits: [2]int{1, 1},
			},
		),
	})
	if err != nil {
		H.s.SendMessage(893577313483096151, err.Error())
	}
}
