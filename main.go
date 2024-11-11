package main

import(
	"log"
	"github.com/coolarif123/Gator/internal/config"
	"github.com/coolarif123/Gator/internal/database"
	"database/sql"
	"os"
)

import _ "github.com/lib/pq"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	
	db, err := sql.Open("postgres", cfg.DB_Url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Create queries
	dbQueries := database.New(db)

	//THIS IS HOW YOU ASSIGN THE CFG TO state.cfg
	//We assign the address for the config pointer to the address of cfg
	//makes a lot of sense once u actually see how to do it
	programState := &State{
		Cfg: &cfg,
		Db: dbQueries,
	}

	commands := Commands{
		Callback: make(map[string]func(*State, Command) error),
	}
	commands.Register("login", Login)
	commands.Register("register", Register)
	commands.Register("reset", Reset)
	commands.Register("users", Users)
	commands.Register("agg", Agg)
	commands.Register("addfeed", middlewareLoggedIn(AddFeed))
	commands.Register("feeds", ListFeeds)
	commands.Register("following", middlewareLoggedIn(Following))
	commands.Register("follow", middlewareLoggedIn(Follow))
	commands.Register("unfollow", middlewareLoggedIn(Unfollow))
	commands.Register("browse", Browse)

	if len(os.Args) < 2 {
		log.Fatal("No command was entered")
		return
	}	

	command := Command{
		Name: os.Args[1],
		Args: os.Args,
	}
	
	err = commands.Run(programState, command)
	if err != nil {
		log.Fatal(err)
	}

}