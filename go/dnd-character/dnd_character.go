package dndcharacter

import (
	"math"
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	// panic("Please implement the Modifier() function")
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	// panic("Please implement the Ability() function")
	// Set the seed for random number generation
	// rand.Seed(time.Now().UnixNano())

	// Define the range (3 to 18)
	min := 3
	max := 18

	// Generate a random number within the specified range
	randomNumber := rand.Intn(max-min+1) + min

	return randomNumber
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	// panic("Please implement the GenerateCharacter() function")

	constitution :=Ability()
	return Character{
		Strength: Ability(),
		Dexterity: Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom: Ability(),
		Charisma: Ability(),
		Hitpoints: 10 + Modifier(constitution),
	}
}
