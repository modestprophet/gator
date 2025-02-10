package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/modestprophet/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		l, err := strconv.Atoi(cmd.args[0])
		if err == nil && l > 0 {
			limit = l
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name:  user.Name,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("get posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found")
		return nil
	}

	for _, post := range posts {
		fmt.Printf("[%s] %s\n%s\n%s\n\n",
			post.PublishedAt.Time.Format("2006-01-02 15:04"),
			post.Title,
			post.Url,
			post.Description.String)
	}
	return nil
}
