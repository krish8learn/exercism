package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	// panic("Please implement the Valid function")

	//trim space
	id = strings.ReplaceAll(id, " ", "")
	if len(id) == 1 {
		return false
	}
	//check the input is valid or not after removing spaces
	if _, err := strconv.Atoi(id); err != nil {
		return false
	}
	sum := 0
	isSecond := false
	for index := len(id) - 1; index >= 0; index-- {
		intValue, _ := strconv.Atoi(string(id[index]))
		if isSecond {
			doubleValue := intValue * 2
			if doubleValue > 9 {
				sum += doubleValue - 9
			} else {
				sum += doubleValue
			}
		} else if !isSecond {
			sum += intValue
		}
		isSecond = !isSecond
	}
	// fmt.Println("-->", sum)
	if sum%10 == 0 {
		return true
	}
	return false
}
