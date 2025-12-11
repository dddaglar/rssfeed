package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/denizekindaglar/rssfeed/internal/database"
	"github.com/google/uuid"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2
	if len(cmd.arg) > 1 {
		return fmt.Errorf("<usage: %s <name>", cmd.name)
	}
	if len(cmd.arg) == 1 {
		n, err := strconv.Atoi(cmd.arg[0])
		if err != nil {
			return err
		}
		limit = n
	}
	curUser, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: uuid.NullUUID{UUID: curUser.ID, Valid: true},
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user", err)
	}
	fmt.Printf("Found %d posts for user %v:\n", len(posts), curUser.Name)
	for _, post := range posts {
		fmt.Printf("%s from %v\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("  %v\n", post.Description.String)
		fmt.Printf("Link %s\n", post.Url)
		fmt.Println("=========")
	}
	return nil
}
