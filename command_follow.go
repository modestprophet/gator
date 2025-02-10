package main

import (
	"context"
	"fmt"

	"github.com/modestprophet/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: follow <feed-url>")
	}
	url := cmd.args[0]

	ff, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("create feed follow: %w", err)
	}

	fmt.Printf("Now following %s as %s\n", ff.FeedName, ff.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("get follows: %w", err)
	}

	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: unfollow <feed-url>")
	}
	url := cmd.args[0]

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return fmt.Errorf("delete feed follow: %w", err)
	}

	fmt.Printf("Stopped following %s\n", url)
	return nil
}
