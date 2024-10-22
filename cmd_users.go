package main

import(
	"context"
	"fmt"
	"strings"
	"github.com/coolarif123/Gator/internal/config"
)

func Users(s *State , cmd Command) error {
	userNames, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	if len(userNames) == 0 {
		fmt.Println("There are no users in the database")
		return nil
	}

	cfg, err := config.Read()
	if err != nil {
		return err
	}

	for _, u := range userNames {
		if strings.Compare(cfg.CurrentUserName, u) == 0 {
			fmt.Printf("* %s (current)", u)
		}
		fmt.Printf("* %s\n", u)
	}
	return nil
}