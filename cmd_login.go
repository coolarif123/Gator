package main

import (
	"errors"
	"context"
	"fmt"
)

func Login(s *State, cmd Command) error {
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
	if err != nil {
		return errors.New("Username not registered in the database")
	}

	err = s.Cfg.SetUser(cmd.Args[2])
	if err != nil {
		return err
	}
	fmt.Printf("Logged in as %s\n", cmd.Args[2])

	return nil
}