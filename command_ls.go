package main

import (
	"errors"
	"fmt"
	"os"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s, %s\n", cmd.name, cmd.description)
	}
	return nil
}

func callbackExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Area")
	for _, loc := range resp.Results {
		fmt.Printf("- %s \n", loc.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Area")
	for _, loc := range resp.Results {
		fmt.Printf("- %s \n", loc.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		errors.New("no location area provided")
	}
	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)

	// PokemonsInArea := locationArea.PokemonEncounters
	for _, pokemon := range locationArea.PokemonEncounters {
		cfg.latestExploredPokemons = append(cfg.latestExploredPokemons, pokemon.Pokemon.Name)
		fmt.Printf("- %s \n", pokemon.Pokemon.Name)
	}

	return nil
}

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	found := false
	for _, name := range cfg.latestExploredPokemons {
		if name == pokemonName {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("%s cannot be caught as it was not found in the latest exploration.", pokemonName)
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	pokemonCatchAttemp := CatchProbability(pokemon.BaseExperience)
	fmt.Printf("throwing a Pokeball at %s...\n", pokemon.Name)
	if pokemonCatchAttemp {
		cfg.caughtPokemons[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil

	} else {
		return fmt.Errorf("%s escaped.", pokemon.Name)
	}
}

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemons[pokemonName]
	if !ok {
		return fmt.Errorf("you haven't caught %s\n", pokemonName)
	}
	fmt.Printf("name: %s\n", pokemon.Name)
	fmt.Printf("experience: %v\n", pokemon.BaseExperience)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	return nil
}

func callbackPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("unnecesary details provided")
	}
	fmt.Printf("your Pokedex contains the following pokemon:\n")
	for _, p := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
