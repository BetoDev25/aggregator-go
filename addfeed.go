package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
	"aggregator-go/internal/database"
)

func addFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	ctx := context.Background()

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("no current user, please register or login first")
	}

	currentUserName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(ctx, currentUserName)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: feedName,
		Url: feedURL,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("Feed info:\n")
	fmt.Printf("- ID: %v\n", feed.ID)
	fmt.Printf("- CreatedAt: %v\n", feed.CreatedAt)
	fmt.Printf("- UpdatedAt: %v\n", feed.UpdatedAt)
	fmt.Printf("- Name: %s\n", feed.Name)
	fmt.Printf("- Url: %s\n", feed.Url)
	fmt.Printf("- UserID: %v\n", feed.UserID)

	return nil
}
