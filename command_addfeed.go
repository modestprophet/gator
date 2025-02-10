package main

import (
	"context"
	"fmt"

	"github.com/modestprophet/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	// Get current user
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("get current user: %w", err)
	}

	// Create feed
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("create feed: %w", err)
	}

	fmt.Printf("Created feed:\nID: %s\nName: %s\nURL: %s\nUserID: %s\nCreatedAt: %s\n",
		feed.ID, feed.Name, feed.Url, feed.UserID, feed.CreatedAt)
	return nil
}
