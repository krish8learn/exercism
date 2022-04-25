package pangram

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	// panic("Please implement the IsPangram function")

	//remove special characters
	input = RemoveSpecialCharacters(input)

	//remove underscore   character
	input = strings.ReplaceAll(input, "_", "")

	//convert them to lower case
	input = strings.ToLower(input)

	//remove numbers   
	input = RemoveNumbers(input)

	fmt.Println(input)
	if len(input) < 26 {
		return false
	}

	countMap := make(map[rune]bool)

	for _, value := range input {
		countMap[value] = true
	}

	return len(countMap) == 26
}

func RemoveSpecialCharacters(input string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	retStr := re.ReplaceAllString(input, "")
	return retStr
}

func RemoveNumbers(input string) string {
	re, err := regexp.Compile(`[0-9]`)
	if err != nil {
		log.Fatal(err)
	}
	retStr := re.ReplaceAllString(input, "")
	return retStr
}
