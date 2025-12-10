package main

import (
	"context"
	"fmt"
	"strconv"

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
	posts, err := s.db.GetPostsForUser(context.Background(), uuid.NullUUID{UUID: curUser.ID, Valid: true})
	if err != nil {
		return err
	}
	if limit > len(posts) {
		limit = len(posts)
	}
	fmt.Printf("here is the count of posts %v", limit)
	for i := 0; i < limit; i++ {
		fmt.Printf("%+v\n", posts[i])
	}
	return nil
}
