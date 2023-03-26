package utils

import (

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/joel-samuel-raj/Horikita/constants"
)

func CheckVerified (Roles []discord.RoleID) bool {
	ans := false
	for _, role := range Roles {
		if role.String() == constants.C.GetVerifiedRole() {
			ans = true
			break
		}
	}
	return ans
}

