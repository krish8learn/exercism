package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	// panic("Please implement the Answer function")

	// to find mathematical operations
	operands := []string{
		"plus", "minus", "multiplied", "divided", "cubed",
	}

	//it collect vals and operands from the input
	operations := []string{}
	result := 0
	//remove ? from string
	question = strings.ToLower(question)
	question = strings.ReplaceAll(question, "?", "")

	//extract numbers and operands from input
	words := strings.Split(question, " ")
	for _, word := range words {
		number, err := strconv.Atoi(word)
		if err != nil {
			//word is not number
			for _, operand := range operands {
				if operand == word {
					operations = append(operations, operand)
				}
			}
		}
		if number != 0 {
			operations = append(operations, strconv.Itoa(number))
		}

	}

	//Non math question
	if len(operations) == 0 {
		return 0, false
	}

	//this slice contain number
	numbers := []int{}
	//this slice contains commands
	commands := []string{}
	if number, _ := strconv.Atoi(operations[0]); number == 0 {
		return 0, false
	} else if number, _ := strconv.Atoi(operations[0]); number != 0 && len(operations) == 1 {
		return number, true
	} else {
		for index, value := range operations {
			if val, _ := strconv.Atoi(value); val != 0 {
				if index+1 < len(operations) {
					//we cannot allow 2 numbers consecutively
					if digit, _ := strconv.Atoi(operations[index+1]); digit != 0 {
						return 0, false
					}
				}
				numbers = append(numbers, val)
			} else {
				if index+1 < len(operations) {
					//we cannot allow 2 operands consecutively
					if digit, _ := strconv.Atoi(operations[index+1]); digit == 0 {
						return 0, false
					}
				}

				commands = append(commands, value)
			}
		}
	}

	//solve the equation
	result = calculates(numbers, commands)
	if result == 0 {
		return 0, false
	}

	return result, true
}

func calculates(sliceInt []int, sliceString []string) int {
	result := 0
	if len(sliceInt) == 0 {
		return 0
	} else if len(sliceInt) == 1 {
		return 0
	}
	for i := 0; i < len(sliceString); i++ {
		var firstDigit, secondDigit int
		if i == 0 {
			firstDigit = sliceInt[i] // First op
		} else {
			firstDigit = result // Subsequent op to use result
		}
		if i+1 < len(sliceInt) {
			secondDigit = sliceInt[i+1]
		}

		switch sliceString[i] {
		case "plus":
			result = firstDigit + secondDigit
		case "minus":
			result = firstDigit - secondDigit
		case "multiplied":
			result = firstDigit * secondDigit
		case "divided":
			result = firstDigit / secondDigit
		default: // Operation we don't support
			return 0
		}
	}
	return result
}
