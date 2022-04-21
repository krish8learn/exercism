package isogram

import (
	"log"
	"regexp"
	"strings"
)

func IsIsogram(word string) bool {
	// panic("Please implement the IsIsogram function")

	//remove special characters
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	word = re.ReplaceAllString(word, "")

	//convert them into small characters
	word = strings.ToLower(word)

	//string to character array
	chars := []rune(word)
	
	var found bool = true
	for index1, value1 := range chars{
		for index2 := index1+1; index2<len(chars); index2++{
			if value1 == chars[index2]{
				found = false
				break
			}
		}
	}

	return found
}
