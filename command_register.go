package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/modestprophet/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("name is required")
	}
	name := cmd.args[0]

	// Check for existing user
	_, err := s.db.GetUser(context.Background(), name)
	if err == nil {
		return fmt.Errorf("user already exists")
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("error checking user: %w", err)
	}

	// Create new user
	id := uuid.New()
	now := time.Now()
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Update config
	s.cfg.SetUser(name)

	fmt.Printf("User created successfully:\n%+v\n", user)
	return nil
}
