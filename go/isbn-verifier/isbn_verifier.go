package isbn

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	// panic("Please implement the IsValidISBN function")
	var sum int
	//check input string length atleast 10
	//remove all the hiphen
	str := strings.ReplaceAll(isbn, "-", "")
	//check string codnition by regex
	matched, _ := regexp.Match(`^\d{9}[\d|X]$`, []byte(str))
	if matched {
		//string structure is proper to be valid isbn
		//reverse the string
		reversedString := reverse(str)
		for index, value := range reversedString {
			if unicode.IsNumber(value) {
				//first digit is X
				number, _ := strconv.Atoi(string(value))
				sum += number * (index + 1)
			} else {
				sum += 10 * (index + 1)
			}
		}

		if sum%11 == 0 {
			return true
		}
	}

	return false
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
