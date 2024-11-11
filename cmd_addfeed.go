package main

import(
	"errors"
	"github.com/google/uuid"
	"github.com/coolarif123/Gator/internal/database"
	"time"
	"context"
	"fmt"
)

//We need to add to the user id teh id of the current user

func AddFeed(s *State, cmd Command, user database.User) error {
	if (len(cmd.Args) < 2){
		return errors.New("No command entered")
	}

	if(len(cmd.Args) == 2 ){
		return errors.New("No name or url entered")
	}

	if (len(cmd.Args) == 3) {
		return errors.New("no url entered")
	}

	if(len(cmd.Args) > 4) {
		return errors.New("too many arguments entered")
	}
	
	newFeedID := uuid.New()

	newFeedFollowParams := database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		UserID:		user.ID,
		FeedID:		newFeedID,
	}

	newFeedParams := database.CreateFeedParams{
		ID: 		newFeedID,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
		Name:	   	cmd.Args[2],
		Url:	   	cmd.Args[3],
		UserID:		user.ID,
	}

	_, err := s.Db.CreateFeed(context.Background(), newFeedParams)
	if err != nil {
		return err
	}

	_, err = s.Db.CreateFeedFollow(context.Background(), newFeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println("New Feed Created")
	return nil
} 