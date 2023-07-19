package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	// URLs that you'll need to paginate through location areas.
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type locationArea struct {
	Name string `json:"name"`
}

var currentConfig config

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Displays the names of 20 location areas in the Pokemon world.")
	fmt.Println("mapb: Displays the previous 20 location areas in the Pokemon world.")
	return nil
}

func commandExit() error {
	return errors.New("exit")
}

func commandMap(cfg *config) error {
	return getAndDisplayLocationAreas(cfg.Next, cfg)
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("You are on the first page. Cannot go back.")
		return nil
	}
	return getAndDisplayLocationAreas(cfg.Previous, cfg)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin) // code come bufio
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world.",
			callback:    commandMapb,
		},
	}

	// Initialize config with the initial URL
	cfg := &config{
		Next: "https://pokeapi.co/api/v2/location-area/",
	}

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		// Handle user input using the cmd map
		cmd, exists := commands[input]
		if !exists {
			fmt.Println("Invalid Command")
			continue
		}

		if err := cmd.callback(); err != nil {
			if err.Error() == "exit" {
				break
			}
			fmt.Printf("Error executing command: %s\n", err)
		}
	}
}
