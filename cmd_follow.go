package main

import(
	"context"
	"github.com/google/uuid"
	"github.com/coolarif123/Gator/internal/database"
	"time"
	"fmt"
)

func Follow(s *State, cmd Command, user database.User) error {
	if (len(cmd.Args) == 2){
		fmt.Println("No feed url was input")
	}

	if (len(cmd.Args) > 3){
		fmt.Println("Please only put in one url")
	}


	feedID , err := s.Db.GetFeedIDFromUrl(context.Background(), cmd.Args[2])
	if err != nil {
		return err
	}

	newFeedFollowParams := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		feedID,
	}

	feedFollowAdded, err := s.Db.CreateFeedFollow(context.Background(), newFeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", feedFollowAdded.FeedName)
	fmt.Printf("%s\n", feedFollowAdded.UserName)

	return nil
} 