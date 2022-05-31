package grains

import (
	"fmt"
	"math"
)

const maxInput = 64

func Square(number int) (uint64, error) {
	// panic("Please implement the Square function")
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("Invalid Input")
	}
	grain := math.Pow(2.0, float64(number-1))
	return uint64(grain), nil
}

func Total() uint64 {
	// panic("Please implement the Total function")
	var total uint64
	for i := 1; i <= 64; i++ {
		add, _ := Square(i)
		total +=add 
	}
	return total
}
