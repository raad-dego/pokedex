package main

import (
	"math/rand"
	"time"
)

func CatchProbability(baseExperience int) bool {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Calculate the chance of catching the Pokemon based on its base experience.
	// The catch probability decreases linearly with higher base experience.
	chance := 100 - baseExperience/10
	if chance < 10 {
		chance = 10 // Ensure the minimum catch probability is 10%
	}

	// Generate a random number between 0 and 99 (inclusive)
	randomNumber := rand.Intn(100)
	// Check if the random number is less than the catch chance
	return randomNumber < chance
}
