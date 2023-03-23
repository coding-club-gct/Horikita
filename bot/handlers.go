package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/joel-samuel-raj/Horikita/constants"
	"github.com/joel-samuel-raj/Horikita/models"
)

var memberAddHandler = func(c *gateway.GuildMemberAddEvent) {
	DiscordUID := c.Member.User.ID
	url := constants.C.ServerURL
	res, err := http.Get(url + "/api/users?populate[0]=userDetail&&filters[discordUID][$eq]=" + DiscordUID.String())
	if err != nil {
		fmt.Println("Failed to fetch details of joined member", err)
	}
	defer res.Body.Close()
	var users []models.User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		fmt.Println(err)
	}
	user := users[0]
	userSnowflake, _ := discord.ParseSnowflake(user.DiscordUID)
	batchSnowflake, _ := discord.ParseSnowflake(constants.C.GetBatchRole(user.UserDetail.Batch - time.Now().Year()))
	deptSnowflake, _ := discord.ParseSnowflake(constants.C.GetDeptRole(user.UserDetail.Dept))
	genderSnowflake, _ := discord.ParseSnowflake(constants.C.GetGenderRole(user.UserDetail.Gender))
	verifiedSnowflake, _ := discord.ParseSnowflake(constants.C.GetVerifiedRole())
	H.s.ModifyMember(c.GuildID, discord.UserID(userSnowflake), api.ModifyMemberData{
		Nick: &user.UserDetail.Name,
		Roles: &[]discord.RoleID{
			discord.RoleID(deptSnowflake),
			discord.RoleID(batchSnowflake),
			discord.RoleID(genderSnowflake),
			discord.RoleID(verifiedSnowflake),
		},
	})
}

var Handlers = []interface{}{memberAddHandler}
