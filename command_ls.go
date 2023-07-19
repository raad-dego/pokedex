package main

import (
	"fmt"
	"os"
)

func callbackHelp() error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s, %s\n", cmd.name, cmd.description)
	}
	return nil
}

func callbackExit() error {
	os.Exit(0)
	return nil
}

// func callbackMap(cfg *config) error {
// 	// return getAndDisplayLocationAreas(cfg.Next, cfg)
// 	return nil
// }

// func callbackMapb(cfg *config) error {
// 	if cfg.Previous == "" {
// 		fmt.Println("You are on the first page. Cannot go back.")
// 		return nil
// 	}
// 	// return getAndDisplayLocationAreas(cfg.Previous, cfg)
// 	return nil
// }
