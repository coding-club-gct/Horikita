package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/joel-samuel-raj/Horikita/constants"
	"github.com/joel-samuel-raj/Horikita/models"
)

func EventRegisterButtonInteraction(ev discord.InteractionEvent, eventID int) *api.InteractionResponse {
	// bearer := "Bearer" + constants.C.ServerApiToken
	url := constants.C.ServerURL + "/api/events/" + strconv.Itoa(eventID) + "?populate[0]=teams"
	res, err := http.Get(url)
	var event models.Event
	err = json.NewDecoder(res.Body).Decode(&event)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(event.Data.ID)
	defer res.Body.Close()
	resp := &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags:   api.EphemeralResponse,
			Content: option.NewNullableString("Please select your team members"),
			Components: discord.ComponentsPtr(
				&discord.UserSelectComponent{
					CustomID: "EventUserSelectComponent",
					Placeholder: "Select ...",
					ValueLimits: [2]int {event.Data.Attributes.MinTeamSize, event.Data.Attributes.MaxTeamSize},
				},
			),
		},
	}
	if err := H.s.RespondInteraction(ev.ID, ev.Token, *resp); err != nil {
		fmt.Println("failed to send interaction callback:", err)
	}
	return resp
}

func EventUserSelectMenuInteraction(ev discord.InteractionEvent, eventID int) *api.InteractionResponse {
	fmt.Println(ev.Data.InteractionType())
	users := ev.Data.(*discord.UserSelectInteraction).Values
	fmt.Println(users)
	resp := &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags: api.EphemeralResponse,
			Content: option.NewNullableString("Team members registration successful"),
		},
	}
	if err := H.s.RespondInteraction(ev.ID, ev.Token, *resp); err != nil {
		fmt.Println("failed to send interaction callback:", err)
	}
	return resp
}
