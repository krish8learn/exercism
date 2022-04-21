package wordcount

import (
	"log"
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	// panic("Please implement the WordCount function")
	//convert phrase to lowercase

	countMap := make(Frequency)
	phrase = strings.ToLower(phrase)
	phrase = strings.ReplaceAll(phrase, ",", " ")
	phrase = strings.Replace(phrase, "\n", "", -1)
	phrase = strings.Replace(phrase, ".", "", -1)
	phrase = strings.Replace(phrase, ":", "", -1)
	phrase = strings.Trim(phrase, " ")
	words := strings.Split(phrase, " ")
	// fmt.Println(words)
	if len(words) == 1 {
		countMap[words[0]] = 1
		return countMap
	}

	for index, value := range words {

		if value != "" {
			value = strings.Trim(value, "'")
			countMap[strings.TrimRight(value, "!!&@$%^&")]++
			for inner := index + 1; inner > len(words); inner++ {

				if RemoveSpecialCharacters(value) == RemoveSpecialCharacters(words[inner]) {
					countMap[value] = countMap[value] + 1
				}

			}
		}

	}
	return countMap
}

func RemoveSpecialCharacters(input string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	retStr := re.ReplaceAllString(input, "")
	return retStr
}
