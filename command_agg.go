package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/modestprophet/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("missing time_between_reqs argument")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	// Handle Ctrl+C
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Immediate first run
	if err := scrapeFeeds(s); err != nil {
		fmt.Printf("Error scraping feeds: %v\n", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := scrapeFeeds(s); err != nil {
				fmt.Printf("Error scraping feeds: %v\n", err)
			}
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("get next feed: %w", err)
	}

	// Mark as fetched before processing to prevent concurrent fetches
	if err := s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return fmt.Errorf("mark feed fetched: %w", err)
	}

	// Fetch and process feed
	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("fetch feed: %w", err)
	}

	for _, item := range rssFeed.Channel.Items {
		fmt.Printf("New post: %s\n", item.Title)
	}

	for _, item := range rssFeed.Channel.Items {
		publishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			// Try alternative format if needed
			publishedAt = time.Now()
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: publishedAt, Valid: true},
			FeedID:      feed.ID,
		})

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("Failed to save post: %v\n", err)
		}
	}

	return nil
}
