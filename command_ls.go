package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s, %s\n", cmd.name, cmd.description)
	}
	return nil
}

func callbackExit(cfg *config) error {
	os.Exit(0)
	return nil
}

var nextURL string // Global variable to store the Next URL for pagination

func callbackMap(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Area")
	for _, loc := range resp.Results {
		fmt.Printf("- %s \n", loc.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Area")
	for _, loc := range resp.Results {
		fmt.Printf("- %s \n", loc.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil

}
