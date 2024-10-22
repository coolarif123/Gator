package main

import(
	"context"
	"github.com/google/uuid"
	"github.com/coolarif123/Gator/internal/database"
	"time"
	"errors"
	"log"
	"fmt"
)

func Register(s *State , cmd Command) error {
	if (len(cmd.Args) < 2){
		return errors.New("No command entered")
	}

	if(len(cmd.Args) == 2 ){
		return errors.New("No username entered")
	}

	if (len(cmd.Args) > 3) {
		return errors.New("Username can only be 1 word!")
	}

	_, err := s.Db.GetUser(context.Background(), cmd.Args[2])
	if err == nil {
		return errors.New("Name Already Exists")
	}

	newUserParams := database.CreateUserParams{
		ID: 		uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name:		cmd.Args[2],
	}

	user, err := s.Db.CreateUser(context.Background(), newUserParams)
	if err != nil {
		return err
	}

	err = s.Cfg.SetUser(cmd.Args[2])
	if err != nil {
		return err
	}

	fmt.Printf("User created: %+v\n", user)
	log.Printf("User created: %+v", user)
	return nil
}