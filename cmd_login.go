package main

import (
	"errors"
	"context"
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
	return nil
}