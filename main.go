package main

import (
	"pokedex/pokeapi"
	"time"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	latestExploredPokemons  []string
	caughtPokemons          map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:          pokeapi.NewClient(time.Hour),
		latestExploredPokemons: make([]string, 0),
		caughtPokemons:         make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
