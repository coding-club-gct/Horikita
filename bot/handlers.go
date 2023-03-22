package bot

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	userSnowflake, err := discord.ParseSnowflake(user.DiscordUID)
	batchSnowflake, err := discord.ParseSnowflake(constants.C.GetBatchRole(user.UserDetail.Batch))
	deptSnowflake, err := discord.ParseSnowflake(constants.C.GetBatchRole(user.UserDetail.Batch))
	genderSnowflake, err := discord.ParseSnowflake(constants.C.GetBatchRole(user.UserDetail.Batch))
	H.s.Client.AddRole(c.GuildID, discord.UserID(userSnowflake), discord.RoleID(batchSnowflake), api.AddRoleData{
		AuditLogReason: "",
	})
	H.s.Client.AddRole(c.GuildID, discord.UserID(deptSnowflake), discord.RoleID(batchSnowflake), api.AddRoleData{
		AuditLogReason: "",
	})
	H.s.Client.AddRole(c.GuildID, discord.UserID(genderSnowflake), discord.RoleID(batchSnowflake), api.AddRoleData{
		AuditLogReason: "",
	})
}

var Handlers = []interface{}{memberAddHandler}
