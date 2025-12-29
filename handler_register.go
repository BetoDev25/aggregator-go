package main

import (
	"fmt"
	"context"
	"database/sql"
	"time"
	"errors"
	"log"

	"github.com/google/uuid"
	"aggregator-go/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, username)
	if err == nil {
		return fmt.Errorf("user %s already exists", username)
	}
	if errors.Is(err, sql.ErrNoRows) {

	} else {
		return fmt.Errorf("error checking username: %w", err)
	}

	insertedUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: username,
	})
	if err != nil {
		return fmt.Errorf("could not insert user: %w", err)
	}

	err = s.cfg.SetUser(insertedUser.Name)
	if err != nil {
		return fmt.Errorf("could not change config user: %w", err)
	}

	log.Printf("successfully created user %s\n", insertedUser.Name)

	return nil
}
