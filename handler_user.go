package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
        if len(cmd.Args) == 0 {
                return fmt.Errorf("no argument provided")
        }

        username := cmd.Args[0]

        err := s.cfg.SetUser(username)
        if err != nil {
                return fmt.Errorf("Could not change username: %w", err)
        }

        fmt.Printf("Username set to %s\n", username)

        return nil
}
