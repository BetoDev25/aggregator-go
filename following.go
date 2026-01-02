package main

import (
	"fmt"
	"context"

	"aggregator-go/internal/database"
)

func following(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not retrieve follows for user: %w", err)
	}

	for _, feed := range follows {
		fmt.Printf("- %s\n", feed.FeedName)
	}

	return nil
}
