package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("username required")
	}

	name := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user does not exist")
		}
		return fmt.Errorf("error getting user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting username in config: %w", err)
	}

	fmt.Printf("User set to: %s\n", cmd.args[0])
	return nil
}
