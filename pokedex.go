package main

type Pokemon struct {
	Name           string
	BaseExperience int
}

var pokedex map[string]Pokemon

func init() {
	pokedex = make(map[string]Pokemon)
}

func CatchPokemon(name string, baseExperience int) *Pokemon {
	return &Pokemon{
		Name:           name,
		BaseExperience: baseExperience,
	}
}
