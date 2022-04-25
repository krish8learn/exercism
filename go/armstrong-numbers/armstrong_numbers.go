package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	// panic("Please implement the IsNumber function")

	inputString := strconv.Itoa(n)
	length := len(inputString)
	sum := 0
	for _, value := range inputString {
		inputInt, _ := strconv.Atoi(string(value))
		sum += int(math.Pow(float64(inputInt), float64(length)))
	}
	if sum == n {
		return true
	}

	return false
}
