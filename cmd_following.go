package main

import(
	"context"
	"fmt"
	"github.com/coolarif123/Gator/internal/database"
)

func Following(s *State, cmd Command, user database.User) error {

	feedFollowSliceForUser, err := s.Db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, f := range feedFollowSliceForUser {
		fmt.Printf("%s\n", f.FeedName)
	}

	return nil
}