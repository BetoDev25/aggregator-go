package main

import (
	"fmt"
	"log"
	"os"

	"aggregator-go/internal/config"
)

type state struct {
	cfg *config.Config
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

	s := &state{
		cfg: &cfg,
	}

	//create commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

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
