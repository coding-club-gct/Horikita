package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"github.com/joel-samuel-raj/Horikita/constants"
	"github.com/joel-samuel-raj/Horikita/models"
	"github.com/joel-samuel-raj/Horikita/types"
	"github.com/joel-samuel-raj/Horikita/utils"
)

type reqPayload struct {
	Data models.Team `json:"data"`
}

func (H *Handler) HandleInteraction(ev *discord.InteractionEvent) *api.InteractionResponse {
	if utils.CheckVerified(ev.Member.RoleIDs) {
		switch data := ev.Data.(type) {
		case *discord.CommandInteraction:
			switch data.Name {
			case "team-create":
				return CreateTeamInteraction(ev, data)
			default:
				return utils.ErrorResponse(fmt.Errorf("unknown command %q", data.Name))
			}
		case discord.ComponentInteraction:
			var payload types.LoadedCustomId
			json.Unmarshal([]byte(data.ID()), &payload)
			switch payload.CustomID {
			case "TeamCreateSelectMenu":
				return CreateTeamSelectMenuInteraction(ev)
			case "TeamCreateMemberSelectMenu":
				return teamMemberSelectInteraction(ev)
			}
		default:
			return utils.ErrorResponse(fmt.Errorf("unknown interaction %T", ev.Data))
		}
	} else {
		return &api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Flags:   api.EphemeralResponse,
				Content: option.NewNullableString("You need to be verified to use this command. For furthur information, contact staffs."),
			},
		}
	}
	return nil
}
func CreateTeamInteraction(ev *discord.InteractionEvent, data *discord.CommandInteraction) *api.InteractionResponse {
	var options struct {
		Arg string `discord:"name"`
	}
	res, err := http.Get(constants.C.ServerURL + "/api/events?filters[open][$eq]=true")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	var events struct {
		Data []models.Event `json:"data"`
	}
	err = json.NewDecoder(res.Body).Decode(&events)
	if err != nil {
		fmt.Println(err)
	}
	if err := data.Options.Unmarshal(&options); err != nil {
		return utils.ErrorResponse(err)
	}
	payload := types.LoadedCustomId{
		CustomID: "TeamCreateSelectMenu",
		Payload:  options.Arg,
	}
	strPayload, _ := json.Marshal(payload)
	opts := make([]discord.SelectOption, 0, len(events.Data))
	for _, event := range events.Data {
		opts = append(opts, discord.SelectOption{
			Label:       event.Attributes.Name,
			Value:       strconv.Itoa(event.ID),
			Description: event.Attributes.ShortLiner,
		})
	}
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags:   api.EphemeralResponse,
			Content: option.NewNullableString("Please select the Event for which you are creating the team"),
			Components: discord.ComponentsPtr(
				&discord.StringSelectComponent{
					CustomID: discord.ComponentID(strPayload),
					Options:  opts,
				},
			),
		},
	}
}

func CreateTeamSelectMenuInteraction(ev *discord.InteractionEvent) *api.InteractionResponse {
	url := constants.C.ServerURL + "/api/teams?populate[0]=event"
	data, _ := ev.Data.(*discord.SelectInteraction)
	var payload types.LoadedCustomId
	json.Unmarshal([]byte(data.CustomID), &payload)
	eventID, _ := strconv.Atoi(data.Values[0])
	_UserID, err := GetUserIDByDiscordUserUID(ev.Member.User.ID.String())
	UserID, err := strconv.Atoi(_UserID)
	jsonData, _ := json.Marshal(reqPayload{
		Data: models.Team{
			EventID:    eventID,
			Name:       payload.Payload,
			TeamLeader: UserID,
		},
	})
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	var TeamResp struct {
		Data struct {
			ID int `json:"id"`
			Attributes struct {
				Event struct {
					Data struct {
						models.Event 
					} `json:"data"`
				}  `json:"event"`
			} `json:"attributes"`
		} `json:"data"`
	}
	json.NewDecoder(res.Body).Decode(&TeamResp)
	_MemberSelectCustomPayload := types.LoadedCustomId {
		CustomID: "TeamCreateMemberSelectMenu",
		Payload: strconv.Itoa(TeamResp.Data.ID),
	}
	MemberSelectCustomPayload, _ := json.Marshal(_MemberSelectCustomPayload)
	if res.StatusCode == 200 {
		return &api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Flags:   api.EphemeralResponse,
				Content: option.NewNullableString("Select the team members for your team including yourself. You will be automatically assigned as team leader for your team"),
				Components: discord.ComponentsPtr(
					&discord.UserSelectComponent{
						CustomID: discord.ComponentID(MemberSelectCustomPayload),
						ValueLimits: [2]int{TeamResp.Data.Attributes.Event.Data.Attributes.MinTeamSize, TeamResp.Data.Attributes.Event.Data.Attributes.MaxTeamSize},
					},
				),
			},
		}
	} else {
		return nil
	}
}

func teamMemberSelectInteraction (ev *discord.InteractionEvent) *api.InteractionResponse {
	data, _ := ev.Data.(*discord.UserSelectInteraction)
	var payload types.LoadedCustomId
	json.Unmarshal([]byte(data.CustomID), &payload)
	url := constants.C.ServerURL + "/api/teams/" + payload.Payload 
	UserIDs := []int{}
	fmt.Println("here")
	for _, userID := range data.Values {
		member, _ := H.s.Member(ev.GuildID, userID)
		verified := false
		for _, role := range member.RoleIDs {
			if role.String() == constants.C.GetVerifiedRole() {
				verified = true
				break
			}
		}
		_userID, _ := GetUserIDByDiscordUserUID(userID.String())
		__userID, _ := strconv.Atoi(_userID)
		UserIDs = append(UserIDs, __userID)
		if !verified {
			return nil
		} 
	}
	var requestPayload struct {
		Data struct {
			Members []int `json:"members"`
		} `json:"data"`
	}
	requestPayload.Data.Members = UserIDs
	jsonRequestPayload, err := json.Marshal(requestPayload)
	fmt.Println(string(jsonRequestPayload))
	if err != nil {
		fmt.Println(err)
	} 
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonRequestPayload))
	if err != nil {
        panic(err)
    }
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
	if resp.StatusCode == 200 {
		return &api.InteractionResponse {
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Content: option.NewNullableString("Team created successfully"),
			},
		}
	} else {
		return nil
	}
}
