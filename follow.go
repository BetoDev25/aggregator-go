package main

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"aggregator-go/internal/database"

)

func follow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1  {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	ctx := context.Background()

	url := cmd.Args[0]

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
	fmt.Printf("- Current user: %s\n", s.cfg.CurrentUserName)

	return nil
}



