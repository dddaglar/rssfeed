package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/denizekindaglar/rssfeed/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arg) != 2 {
		return fmt.Errorf("<usage: %s <name> <url>", cmd.name)
	}
	name, url := cmd.arg[0], cmd.arg[1]
	newFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Name:      sql.NullString{String: name, Valid: true},
		Url:       url,
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to add feed: %v", err)
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: newFeed.CreatedAt,
		UpdatedAt: newFeed.CreatedAt,
		UserID:    newFeed.UserID,
		FeedID:    uuid.NullUUID{UUID: newFeed.ID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %v", err)
	}

	fmt.Printf("%+v\n", newFeed)
	return nil
}
