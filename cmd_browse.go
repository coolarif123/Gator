package main

import(
	"fmt"
	"context"
	"strconv"
	"errors"
)

func Browse(s *State, cmd Command) error {
	var limit int32 = 2

	if (len(cmd.Args) == 3){
		arg, err := strconv.ParseInt(cmd.Args[2], 10, 32)
		if err != nil {
			return err
		}

		limit = int32(arg)
	}

	if (len(cmd.Args) > 3 ){
		return errors.New("too many arguments input")
	}

	posts , err := s.Db.GetPostsForUser(context.Background(), limit)
	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Printf("Title: %s\n", p.Title.String)
		fmt.Printf("URL: %s\n", p.Url)
		fmt.Printf("Description: %s\n", p.Description.String)
		fmt.Printf("Published At: %s\n", p.PublishedAt)
		fmt.Println("")
		fmt.Println("--------------------------------------")
	}

	return nil
}