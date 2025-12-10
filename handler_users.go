package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %v", err)
	}
	for _, user := range users {
		if s.config.CurrentUserName == user {
			fmt.Printf("* %s (current)\n", user)
			continue
		}
		fmt.Printf("* %s\n", user)
	}
	return nil

}
