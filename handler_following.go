package main

import (
	"context"
	"fmt"

	"github.com/denizekindaglar/rssfeed/internal/database"
	"github.com/google/uuid"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), uuid.NullUUID{UUID: user.ID, Valid: true})
	if err != nil {
		return err
	}
	for _, feedFollow := range feedFollows {
		if feedFollow.UserName == user.Name {
			fmt.Printf("* %v (%s) \n", feedFollow.FeedName.String, user.Name)
		}
	}
	return nil
}
