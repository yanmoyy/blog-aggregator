package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yanmoyy/blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed by url: %w", err)
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user by name: %w", err)
	}
	row, err := s.db.CreatFeedFollow(context.Background(), database.CreatFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't creat feed_follow: %w", err)
	}
	fmt.Println("Follow feed successfully:")
	printCreatFeedFollowRow(row)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get user by name: %w", err)
	}
	rows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follow list by userId: %w", err)
	}
	fmt.Printf("=== User (%s)'s following feeds ===\n", user.Name)
	for i, row := range rows {
		fmt.Printf("Feed %d\n", i)
		printGetFeedFollowForUserRow(row)
	}
	return nil
}

func printCreatFeedFollowRow(row database.CreatFeedFollowRow) {
	fmt.Printf("* ID:            %s\n", row.ID)
	fmt.Printf("* Created:       %v\n", row.CreatedAt)
	fmt.Printf("* Updated:       %v\n", row.UpdatedAt)
	fmt.Printf("* UserID:        %s\n", row.UserID)
	fmt.Printf("* FeedID:        %s\n", row.FeedID)
	fmt.Printf("* UserName:      %s\n", row.UserName)
	fmt.Printf("* FeedName:      %s\n", row.FeedName)
}

func printGetFeedFollowForUserRow(row database.GetFeedFollowsForUserRow) {
	fmt.Printf("* ID:            %s\n", row.ID)
	fmt.Printf("* Created:       %v\n", row.CreatedAt)
	fmt.Printf("* Updated:       %v\n", row.UpdatedAt)
	fmt.Printf("* UserID:        %s\n", row.UserID)
	fmt.Printf("* FeedID:        %s\n", row.FeedID)
	fmt.Printf("* FeedName:      %s\n", row.FeedName)
}

func printFeedFollow(ff database.FeedFollow) {
	fmt.Printf("* ID:            %s\n", ff.ID)
	fmt.Printf("* Created:       %v\n", ff.CreatedAt)
	fmt.Printf("* Updated:       %v\n", ff.UpdatedAt)
	fmt.Printf("* UserID:        %s\n", ff.UserID)
	fmt.Printf("* FeedID:        %s\n", ff.FeedID)
}
