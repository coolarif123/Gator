package main

import(
	"errors"
	"context"
	"fmt"
)

func Reset(s *State , cmd Command) error {
	if len(cmd.Args) != 2 {
		errors.New("For reset only enter the word reset and no other arguments!")
	}

	err := s.Db.Reset(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Database has been reset")
	return nil
}