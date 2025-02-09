package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("get users: %w", err)
	}

	currentName := s.cfg.CurrentUserName
	for _, user := range users {
		marker := ""
		if user.Name == currentName {
			marker = " (current)"
		}
		fmt.Printf("* %s%s\n", user.Name, marker)
	}
	return nil
}
