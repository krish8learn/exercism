package allyourbase

import (
	"fmt"
	"math"
)

const charsetNumbers = "0123456789ABCDEF"

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	// panic("Please implement the ConvertToBase .function")

	// base validation error
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	} else if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}

	// digits validation error
	for _, value := range inputDigits {
		if value < 0 || value > inputBase-1 {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
	}

	//check empty slice
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	//reverse the inputDigits
	intReverse(inputDigits)

	//convert decimal
	decimalNumber := otherToDecimal(inputBase, inputDigits)

	//convert decimal to required format
	outputBaseNumber := decimalToOther(outputBase, decimalNumber)

	//reverse the outputDigits
	intReverse(outputBaseNumber)

	return outputBaseNumber, nil

}

func otherToDecimal(base int, digits []int) (decimal int) {

	for i := len(digits) - 1; i >= 0; i-- {
		decimal += int(math.Pow(float64(base), float64(i))) * digits[i]
	}

	return decimal
}

func decimalToOther(base int, decimalNumber int) []int {

	var outputNumber []int
	if decimalNumber == 0 {
		outputNumber = append(outputNumber, 0)
		return outputNumber
	}
	for decimalNumber > 0 {
		remainder := decimalNumber % base
		outputNumber = append(outputNumber, remainder)
		decimalNumber /= base
	}
	return outputNumber
}

func intReverse(input []int) {
	inputLen := len(input)
	inputMid := inputLen / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		input[i], input[j] = input[j], input[i]
	}
}
