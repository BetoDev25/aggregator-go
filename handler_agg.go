package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	/*
	if len(cmd.Args) != 1 {
		return fmt.Errorf("no argument provided")
	}
	*/

	ctx := context.Background()

	//URL := cmd.Args[0]

	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("could not get feed: %w", err)
	}

	/*
	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Link: %s\n", feed.Channel.Link)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		fmt.Printf("Item %d:\n", i)
		fmt.Printf("- Title: %s\n", item.Title)
		fmt.Printf("- Link: %s\n", item.Link)
		fmt.Printf("- Description: %s\n", item.Description)
		fmt.Printf("- PubDate: %s\n", item.PubDate)
	}
	*/

	fmt.Printf("%+v\n", feed)

	return nil
}
