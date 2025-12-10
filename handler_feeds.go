package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID.UUID)
		if err != nil {
			return err
		}
		fmt.Printf("* %v %s (%s)\n", feed.Name.String, feed.Url, user.Name)
	}
	return nil
}
