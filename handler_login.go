package main

import (
	"fmt"
	"context"
	"database/sql"
	"errors"
)

func handlerLogin(s *state, cmd command) error {
        if len(cmd.Args) == 0 {
                return fmt.Errorf("no argument provided")
        }

        username := cmd.Args[0]

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user %s does not exist", username)
	}
	if err != nil {
		return fmt.Errorf("error looking up user: %w", err)
	}

        err = s.cfg.SetUser(username)
        if err != nil {
                return fmt.Errorf("Could not change username: %w", err)
        }

        fmt.Printf("Username set to %s\n", username)

        return nil
}
