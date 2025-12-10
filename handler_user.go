package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arg) != 1 {
		return fmt.Errorf("<usage: %s <name>", cmd.name)
	}
	username := cmd.arg[0]

	//check if user exists, if not return error
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("failed %v", err)
	}
	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set user %w", err)
	}
	fmt.Println("user has been set successfully")
	return nil
}
