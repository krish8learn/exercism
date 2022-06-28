package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {

	output := ""
	if input == "" {
		return output
	}
	char := ""
	count := 0
	//will identify each character
	for i, value := range input {
		if i == 0 {
			// for first character storing to compare with next one
			char = string(value)
			count++
		} else if char != string(value) {
			// stored character not equal to the current character
			//add to the output string
			output += fmt.Sprintf("%d%s", count, char)
			//store new character
			char = string(value)
			count = 1
		} else if char == string(value) {
			//stored character and current character same
			count++
		}

		if i == len(input)-1 {
			//last iteration nothing to compare
			output += fmt.Sprintf("%d%s", count, char)
		}
	}

	//replace all the single 1[coming before a letter] in string
	var withoutOne string
	for index, val := range output {
		if string(val) == "1" {
			if _, err := strconv.Atoi(string(output[index+1])); err != nil {
				//cannot be conveted to integer
				continue
			}
		}
		withoutOne += string(val)
	}

	return withoutOne
}

func RunLengthDecode(input string) string {
	var decoded string

	for len(input) > 0 {
		/*
			it will return count of digit present in string untill character occurred
			no digit in string, return 0
			only digit in string, return -1
		*/
		i := strings.IndexFunc(input, func(r rune) bool {

			return !unicode.IsDigit(r)

		})

		n := 1

		if i != 0 {

			n, _ = strconv.Atoi(input[:i])

		}

		decoded += strings.Repeat(string(input[i]), n)

		input = input[i+1:]

	}

	return decoded
}
