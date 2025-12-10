package main

import (
	"context"
	"fmt"

	"github.com/denizekindaglar/rssfeed/internal/database"
	"github.com/google/uuid"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arg) != 1 {
		return fmt.Errorf("<usage: %s <feed_url>", cmd.name)
	}
	feedURL := cmd.arg[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: uuid.NullUUID{UUID: user.ID, Valid: true},
		FeedID: uuid.NullUUID{UUID: feed.ID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %v", err)
	}
	return nil

}
