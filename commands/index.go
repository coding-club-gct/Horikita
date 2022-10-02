package commands

import (
	"github.com/diamondburned/arikawa/v3/session"
)

func Init(s *session.Session) {
	functions := [...]func(s *session.Session){Ping}
	for _, command := range functions {
		command(s)
	}
}
