package main

import (
	"fmt"
	"context"
	"aggregator-go/internal/database"
)

func unfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not retrieve feed: %w", err)
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), params)

	return err
}
