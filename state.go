package main

import(
	"errors"
	"github.com/coolarif123/Gator/internal/config"
	"github.com/coolarif123/Gator/internal/database"
	"context"	
)

type State struct {
	Db 	*database.Queries
	Cfg *config.Config
}

type Command struct {
	Name		string
	Args		[]string  
}

type Commands struct {
	Callback 	map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error){
	c.Callback[name] = f
	return
}

func (c *Commands) Run(s *State, cmd Command) error {
	if f , exists := c.Callback[cmd.Name]; !exists {
		return errors.New("command that was ran does not exist")
	} else {
		err := f(s, cmd)
		if err != nil {
			return err
		}
		return nil
	}	
}

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
        // Retrieve user details
        user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
        if err != nil {
            return err
        }

        // Call the original handler, passing the retrieved user
        return handler(s, cmd, user)
    }	
}

