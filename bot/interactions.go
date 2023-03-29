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
	"github.com/joel-samuel-raj/Horikita/httpUtil"
	"github.com/joel-samuel-raj/Horikita/models"
	"github.com/joel-samuel-raj/Horikita/types"
	"github.com/joel-samuel-raj/Horikita/utils"
)

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
	client := httpUtil.CreateHTTPClient()
	req, _ := http.NewRequest("GET", constants.C.Strings["SERVER_URL"]+"/api/events?filters[open][$eq]=true", nil)
	httpUtil.AddAuthorizationHeader(req)
	res, err := client.Do(req)
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
	data, _ := ev.Data.(*discord.SelectInteraction)
	var payload types.LoadedCustomId
	json.Unmarshal([]byte(data.CustomID), &payload)
	eventID, _ := strconv.Atoi(data.Values[0])
	jsonData, _ := json.Marshal(nextPayload{
		EventID:   eventID,
		EventName: payload.Payload,
	})
	_MemberSelectCustomPayload := types.LoadedCustomId{
		CustomID: "TeamCreateMemberSelectMenu",
		Payload:  string(jsonData),
	}
	MemberSelectCustomPayload, _ := json.Marshal(_MemberSelectCustomPayload)
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags:   api.EphemeralResponse,
			Content: option.NewNullableString("Select the team members for your team including yourself. You will be automatically assigned as team leader for your team"),
			Components: discord.ComponentsPtr(
				&discord.UserSelectComponent{
					CustomID:    discord.ComponentID(MemberSelectCustomPayload),
					ValueLimits: [2]int{1, 6},
				},
			),
		},
	}
}

func teamMemberSelectInteraction(ev *discord.InteractionEvent) *api.InteractionResponse {
	data, _ := ev.Data.(*discord.UserSelectInteraction)
	var payload types.LoadedCustomId
	json.Unmarshal([]byte(data.CustomID), &payload)
	var event nextPayload
	json.Unmarshal([]byte(payload.Payload), &event)
	url := constants.C.Strings["SERVER_URL"] + "/api/teams"
	UserIDs := []int{}
	for _, userID := range data.Values {
		member, _ := H.s.Member(ev.GuildID, userID)
		verified := false
		for _, role := range member.RoleIDs {
			if role.String() == constants.C.GetVerifiedRole() {
				verified = true
				break
			}
		}
		if !verified {
			return utils.SendResponse(api.EphemeralResponse, "Oops! Some of your team members are not verified")
		}
		_userID, _ := GetUserIDByDiscordUserUID(userID.String())
		__userID, _ := strconv.Atoi(_userID)
		UserIDs = append(UserIDs, __userID)
	}
	type RequestPayload struct {
		EventID    int    `json:"event"`
		TeamLeader int    `json:"teamLeader"`
		Name       string `json:"name"`
		Members    []int  `json:"members"`
	}
	var reqPayload struct {
		Data RequestPayload `json:"data"`
	}
	_teamLeaderUserID, _ := GetUserIDByDiscordUserUID(ev.Member.User.ID.String())
	teamLeaderUserID, _ := strconv.Atoi(_teamLeaderUserID)
	reqPayload.Data = RequestPayload{
		EventID:    event.EventID,
		TeamLeader: teamLeaderUserID, 
		Name:       event.EventName,
		Members:    UserIDs,
	}
	jsonRequestPayload, err := json.Marshal(reqPayload)
	if err != nil {
		fmt.Println(err)
	}
	client := httpUtil.CreateHTTPClient()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestPayload))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	httpUtil.AddAuthorizationHeader(req)
	client.Do(req)
	return &api.InteractionResponse{
		Type: api.MessageInteractionWithSource,
		Data: &api.InteractionResponseData{
			Flags:   api.EphemeralResponse,
			Content: option.NewNullableString("Team created successfully"),
		},
	}
}

type nextPayload struct {
	EventID   int
	EventName string
}
