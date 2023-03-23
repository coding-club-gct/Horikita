package bot

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joel-samuel-raj/Horikita/constants"
	"github.com/pkg/errors"
)

func overwriteCommands(s *state.State) error {
	app, err := s.CurrentApplication()
	if err != nil {
		return errors.Wrap(err, "cannot get current app ID")
	}

	_, err = s.BulkOverwriteCommands(app.ID, Commands)
	return err
}

func GetUserIDByDiscordUserUID(discordUID string) (string, error) {
	res, err := http.Get(constants.C.ServerURL + "/api/users?filters[discordUID][$eq]=" + discordUID)
	if err != nil {
		return "", err
	} else {
		defer res.Body.Close()
		var UserIDResps []struct {
			ID int `json:"id"`
		}
		err = json.NewDecoder(res.Body).Decode(&UserIDResps)
		if err != nil {
			return "", err
		} else {
			UserID := strconv.Itoa(UserIDResps[0].ID)
			return UserID, nil
		}
	}
}
