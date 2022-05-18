package chance

import (
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	// panic("Please implement the SeedWithTime function")
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	// panic("Please implement the RollADie function")
	min := 1
	max := 20
	
	return rand.Intn(max-min) + min
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	// panic("Please implement the GenerateWandEnergy function")
	min := 0
	max := 12
	value := float64(rand.Intn(max-min) + min)
	return value
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	// panic("Please implement the ShuffleAnimals function")
	elem := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(elem), func(i, j int) {
		elem[i], elem[j] = elem[j], elem[i]
	})
	return elem
}
