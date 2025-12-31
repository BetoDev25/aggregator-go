package main

import (
	"fmt"
	"context"
)

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("could not get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("- Name: %s\n", feed.Name)
		fmt.Printf("- URL: %s\n", feed.Url)
		fmt.Printf("- Created by: %s\n", feed.Name_2)
		fmt.Println()
	}

	return nil
}
