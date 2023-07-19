package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world.",
			// callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world.",
			// callback:    callbackMapb,
		},
	}
}
func startRepl() {
	for {
		scanner := bufio.NewScanner(os.Stdin) // code come bufio
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		commandName := cleaned[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		command.callback()
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}