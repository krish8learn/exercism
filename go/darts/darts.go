package darts

import "math"

const (
	inner  = 1.0
	middle = 5.0
	outer  = 10.0
)

func Score(x, y float64) int {
	// panic("Please implement the Score function")
	hypotenuse := math.Hypot(x, y)
	switch {
	case hypotenuse <= inner:
		return 10
	case hypotenuse <= middle:
		return 5
	case hypotenuse <= outer:
		return 1
	default:
		return 0
	}
}
