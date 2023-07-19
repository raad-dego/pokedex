package main

import (
	"fmt"
	"log"
	"pokedex/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient()

	resp, err := pokeClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(resp)
}
