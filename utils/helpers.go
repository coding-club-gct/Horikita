package utils

import (
	"strconv"

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

func StringArrayToIntArray (array []string) []int {
	intArray := []int{}
	for _, i := range array {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        intArray = append(intArray, j)
    }
	return intArray
}