package utils

import (
	"strconv"

	"github.com/diamondburned/arikawa/v3/discord"
)

func CheckVerified (Roles []discord.RoleID) bool {
	ans := false
	for _, role := range Roles {
		if role == 893577313218875467 {
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