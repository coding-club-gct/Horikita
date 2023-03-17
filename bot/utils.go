package bot

import (
	"github.com/diamondburned/arikawa/v3/state"
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