package main

import (
	"context"
	"fmt"

	"github.com/modestprophet/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	// Create feed
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		Name:   name,
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("create feed: %w", err)
	}

	fmt.Printf("Created feed:\nID: %s\nName: %s\nURL: %s\nUserID: %s\nCreatedAt: %s\n",
		feed.ID, feed.Name, feed.Url, feed.UserID, feed.CreatedAt)

	// After creating feed, auto-follow it
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("auto-follow feed: %w", err)
	}
	fmt.Printf("Added feed %s and auto-followed\n", feed.Name)

	return nil
}
