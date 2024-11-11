package main

import(
	"context"
	"fmt"
)

func ListFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetFeed(context.Background())
	if err != nil {
		return err
	}

	for _, f := range feeds {
		fmt.Printf("%s\n", f.Name)
		fmt.Printf("%s\n", f.Url)
		creator, err := s.Db.GetFeedCreator(context.Background(), f.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", creator)
	}

	return nil
}