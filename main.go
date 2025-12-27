package main

import (
	"fmt"
	"log"
	"os"

	"aggregator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	cfg.SetUser("beto")

	fmt.Printf("Current config file:\n %s\n", cfg)
}
