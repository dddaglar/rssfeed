package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetTable(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users table: %v", err)
	}
	fmt.Println("users table has been reset successfully")
	return nil

}
