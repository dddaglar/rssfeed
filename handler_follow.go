package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/denizekindaglar/rssfeed/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arg) != 1 {
		return fmt.Errorf("<usage: %s <feed_url>", cmd.name)
	}
	feedURL := cmd.arg[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to get feed by url: %v", err)
	}
	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
		FeedID: uuid.NullUUID{UUID: feed.ID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}
	fmt.Printf("user %s is now following feed %v\n", user.Name, newFeedFollow.FeedName)
	return nil
}
