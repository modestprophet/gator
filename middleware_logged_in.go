package main

import (
	"context"
	"fmt"

	"github.com/modestprophet/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		//Get current user
		currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("get current user: %w", err)
		}
		return handler(s, cmd, currentUser)
	}
}
