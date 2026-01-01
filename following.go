package main

import (
	"fmt"
	"context"
)

func following(s *state, cmd command) error {
	ctx := context.Background()

	currentUserName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(ctx, currentUserName)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("could not retrieve follows for user: %w", err)
	}

	for _, feed := range follows {
		fmt.Printf("- %s\n", feed.FeedName)
	}

	return nil
}
