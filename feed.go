package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MudassirDev/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	feedName := cmd.Args[0]
	feelUrl := cmd.Args[1]

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

	_, err = s.db.FollowFeed(context.Background(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}

	return nil
}

func handleGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to retrieve feeds: %v", err)
	}
	for _, feed := range feeds {
		fmt.Printf("feed: %v, user: %v\n", feed.Name, feed.Username)
	}
	return nil
}

func handleFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedWithUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("no such field exists: %v", err)
	}
	_, err = s.db.FollowFeed(context.Background(), database.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to follow feed: %v", err)
	}
	return nil
}

func handleFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get following: %v", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func handleUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]

	feed, err := s.db.GetFeedWithUrl(context.Background(), url)

	err = s.db.Unfollow(context.Background(), database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %v", err)
	}

	return nil
}
