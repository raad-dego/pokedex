package main

import (
	"fmt"
	"log"
	"os"
	"pokedex/pokeapi"
)

func callbackHelp() error {
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

var nextURL string // Global variable to store the Next URL for pagination

func callbackMap() error {
	pokeClient := pokeapi.NewClient()

	// if nextURL == "" {
	// 	resp, err := pokeClient.ListLocationAreas()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	printLocationAreas(resp)
	// 	nextURL = resp.Next
	// } else {
	resp, err := pokeClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Area")
	for _, loc := range resp.Results {
		fmt.Printf("- %s \n", loc.Name)
	}

	return nil
}
