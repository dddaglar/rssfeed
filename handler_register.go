package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/denizekindaglar/rssfeed/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arg) != 1 {
		return fmt.Errorf("<usage: %s <name>", cmd.name)
	}
	username := cmd.arg[0]
	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		log.Fatal("failed", err)
	}
	s.config.SetUser(usr.Name)
	fmt.Printf("user %s has been registered and set successfully\n", usr.Name)
	return nil

}
