package romannumerals

import "fmt"

func ToRomanNumeral(input int) (string, error) {
	// panic("Please implement the ToRomanNumeral function")
	//check input limit
	if input < 1 || input > 3000 {
		return "", fmt.Errorf("Input  not in the Roman numerical Limits")
	}
	
	//convert input to numerical elements
	onesInput := input % 10
	tensInput := input / 10 % 10
	hundredsInput := input / 100 % 10
	thousandInput := input / 1000 % 10

	//final 
	output := Thousands(thousandInput) + Hundreds(hundredsInput) + Tens(tensInput) + Ones(onesInput)

	return output, nil
}

func Ones(digit int) string {
	switch digit {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	case 0:
		return ""
	default:
		return "Invalid Digit"
	}
}

func Tens(digit int) string {
	switch digit {
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	case 0:
		return ""
	default:
		return "Invalid digit"
	}
}

func Hundreds(digit int) string {
	switch digit {
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	case 0:
		return ""
	default:
		return "Invalid digit"
	}
}

func Thousands(digit int) string {
	switch digit {
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	case 0:
		return ""
	default:
		return "Invalid digit"
	}
}
