package main

import _ "github.com/lib/pq"

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	"aggregator-go/internal/database"

	"aggregator-go/internal/config"
)

type state struct {
	cfg *config.Config
	db *database.Queries
}

func main() {
	if len(os.Args) < 2 {
		log.Println("no arguments found")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//connect to database
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s := &state{
		cfg: &cfg,
		db: dbQueries,
	}

	//create commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", addFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", follow)
	cmds.register("following", following)

	name := os.Args[1]
	args := os.Args[2:]

	cmd := command{
		Name: name,
		Args: args,
	}

	if err := cmds.run(s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
