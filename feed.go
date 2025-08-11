package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MudassirDev/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feelUrl := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("user not logged in: %v", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       feelUrl,
		Name:      feedName,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create the feed: %v", err)
	}
	fmt.Println(feed)
	return nil
}
