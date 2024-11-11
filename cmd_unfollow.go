package main

import(
	"context"
	"github.com/coolarif123/Gator/internal/database"
	"fmt"
)

func Unfollow(s *State, cmd Command, user database.User) error {
	err := s.Db.DeleteFeedFollow(context.Background(), cmd.Args[2])
	if err != nil {
		return err
	}
	fmt.Printf("%s has been unfollwed by %s\n", cmd.Args[2], user.Name)
	return nil
}