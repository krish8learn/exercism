package squareroot

import (
	"fmt"
	"math"
)

func SquareRoot(number int) (int, error) {
	result := math.Sqrt(float64(number))
	return int(result), nil
}

func OptionalSquareRoot(number int) (int, error) {
	if number <= 0 {
		return 0, fmt.Errorf("cannot compute square root of a negative number: %d", number)
	}

	x := float64(number)
	tolerance := 1e-10
	for {
		root := 0.5 * (x + float64(number)/x)

		diff := x - root
		if diff < 0 {
			diff = -diff
		}

		if diff < tolerance {
			return int(root), nil
		}

		x = root
	}
}
