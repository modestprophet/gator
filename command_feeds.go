package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeedsWithUser(context.Background())
	if err != nil {
		return fmt.Errorf("list feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %s (%s) - added by %s\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}
