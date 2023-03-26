package bot

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/diamondburned/arikawa/v3/state"
	"github.com/joel-samuel-raj/Horikita/constants"
	"github.com/joel-samuel-raj/Horikita/httpUtil"
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
	client := httpUtil.CreateHTTPClientWithBearerToken()
	req, _ := http.NewRequest("GET", constants.C.Strings["SERVER_URL"]+"/api/users?filters[discordUID][$eq]="+ discordUID, nil)
	res, err := client.Do(req)

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
