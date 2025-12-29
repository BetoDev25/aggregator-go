package main

import (
	"fmt"
	"context"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.Reset(ctx)
	if err != nil {
		return fmt.Errorf("could not reset table: %w", err)
	}

	fmt.Println("database successfully reset")

	return nil
}
