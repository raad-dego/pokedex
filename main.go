package main

import (
	"pokedex/pokeapi"
	"time"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	latestExploredPokemons []string

}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		latestExploredPokemons: make([]string, 0),
	}
	startRepl(&cfg)
}
