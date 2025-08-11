package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("failed to fetch the feed: %v", err)
	}

	fmt.Println(*feed)
	return nil
}
