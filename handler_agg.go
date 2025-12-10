package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/denizekindaglar/rssfeed/internal/database"
)

func scrapeFeed(s *state, nextFeed database.GetNextFeedToFetchRow) {
	err := s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
	})
	if err != nil {
		fmt.Printf("failed to mark feed fetched: %v", err)
		return
	}
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		fmt.Printf("failed to fetch feed: %v", err)
		return
	}
	fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	fmt.Printf("Feed Description: %s\n", feed.Channel.Description)
	//save posts to db
	for _, item := range feed.Channel.Item {
		pubAt, _ := time.Parse(time.RFC1123Z, item.PubDate)
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: pubAt, Valid: true},
			FeedID:      uuid.NullUUID{UUID: nextFeed.ID, Valid: true},
		})
		if err != nil {
			fmt.Printf("failed to create post: %v", err)
			continue
		}
	}
}

func scrapeFeeds(s *state) {
	fmt.Println("Scraping feeds...")
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("failed to get next feed to fetch: %v", err)
		return
	}
	scrapeFeed(s, nextFeed)

}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arg) != 1 {
		return fmt.Errorf("<usage: %s <time_between_requests>", cmd.name)
	}
	time_between_reqs, err := time.ParseDuration(cmd.arg[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", time_between_reqs.String())
	ticker := time.NewTicker(time_between_reqs)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
