package main

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"aggregator-go/internal/database"

)

func follow(s *state, cmd command) error {
	if len(cmd.Args) != 1  {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	ctx := context.Background()

	url := cmd.Args[0]

	currentUserName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(ctx, currentUserName)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	feed, err := s.db.GetFeedURL(ctx, url)
	if err != nil {
		return fmt.Errorf("could not get feed: %w", err)
	}

	follow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("- Name: %s\n", follow.FeedName)
	fmt.Printf("- Current user: %s\n", currentUserName)

	return nil
}



