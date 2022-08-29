package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	// panic("Please implement the Atbash function")
	output := ""
	count := 0
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			//get character value in rune
			letter = unicode.ToLower(letter)
			letter = rune('a' + ('z' - letter))
		}
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) {
			//keep adding all characters
			output += string(letter)
			count++
			if count%5 == 0 {
				//keep spaces after 5 characters
				output += " "
			}
		}

	}
	//remove end spaces
	return strings.TrimSpace(output)
}
