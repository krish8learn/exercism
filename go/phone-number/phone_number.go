package phonenumber

import (
	"errors"
	"strconv"
	"strings"
)

// (NXX)-NXX-XXXX

func Number(phoneNumber string) (string, error) {
	// panic("Please implement the Number function")

	//remove spaces
	input := strings.ReplaceAll(phoneNumber, " ", "")
	//remove hyphen
	input = strings.ReplaceAll(input, "-", "")
	//remove brackets
	input = strings.ReplaceAll(input, "(", "")
	input = strings.ReplaceAll(input, ")", "")
	// remove +
	input = strings.ReplaceAll(input, "+", "")
	//remove .
	input = strings.ReplaceAll(input, ".", "")
	//check number validity
	_, err := strconv.Atoi(input)
	if err != nil {
		return "", err
	}

	inputLength := len(input)
	//chcek length
	//10
	if inputLength == 10 {
		zerothElem, _ := strconv.Atoi(string(input[0]))
		thirdElem, _ := strconv.Atoi(string(input[3]))
		if zerothElem < 2 || thirdElem < 2 {
			return "", errors.New("Nth number are not right")
		}
	} else if inputLength == 11 {
		zerothElem, _ := strconv.Atoi(string(input[0]))
		secondElem, _ := strconv.Atoi(string(input[1]))
		thirdElem, _ := strconv.Atoi(string(input[4]))
		if secondElem < 2 || thirdElem < 2 || zerothElem != 1 {
			return "", errors.New("Nth numbers or country code are not right")
		}
	} else if inputLength > 11 || inputLength < 10 {
		return "", errors.New("input length is not right")
	}

	return strings.ReplaceAll(input, "1", ""), nil

}

func AreaCode(phoneNumber string) (string, error) {
	// panic("Please implement the AreaCode function")
	//get the complete number
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	output := ""
	for index := 0; index <= 2; index++ {
		output += string(number[index])
	}
	return output, nil
}

func Format(phoneNumber string) (string, error) {
	// panic("Please implement the Format function")
	//get the complete number
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	output := ""
	for index, value := range number {
		if index == 0 {
			output += "(" + string(value)
			continue
		} else if index == 2 {
			output += string(value) + ")"
			continue
		} else if index == 3 {
			output += " " + string(value)
			continue
		} else if index == 6 {
			output += "-" + string(value)
			continue
		}
		output += string(value)
	}
	// fmt.Println(output)
	return output, nil
}
