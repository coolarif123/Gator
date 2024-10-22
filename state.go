package main

import(
	"errors"
	"github.com/coolarif123/Gator/internal/config"
	"github.com/coolarif123/Gator/internal/database"	
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
