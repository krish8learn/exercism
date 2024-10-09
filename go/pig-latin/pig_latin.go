package piglatin

import "strings"

const (
	vowels = "aeiou"
)

func Sentence(sentence string) string {
	// panic("Please implement the Sentence function")

	words := strings.Split(strings.ToLower(sentence), " ")
	for index, word := range words {
		if ruled1Word := Rule1(word); ruled1Word != "" {
			words[index] = ruled1Word
		} else if ruled3Word := Rule3(word); ruled3Word != "" {
			words[index] = ruled3Word
		} else if ruled2Word := Rule2(word); ruled2Word != "" {
			words[index] = ruled2Word
		} else if ruled4Word := Rule4(word); ruled4Word != "" {
			words[index] = ruled4Word
		}
	}

	result := strings.Join(words, " ")
	return result
}

func Rule1(word string) string {
	// if the word starts with vowels
	firstChar := rune(word[0])
	if strings.ContainsRune(vowels, firstChar) {
		return word + "ay"
	}

	// if the word has prefix like "xr" or "yt"
	if strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "yt") {
		return word + "ay"
	}

	return ""
}

func Rule2(word string) string {
	// if the word starts with consonants
	for index, char := range word {
		if strings.ContainsRune(vowels, char) {
			return word[index:] + word[:index] + "ay"
		}
	}
	return ""
}

func Rule3(word string) string {
	// if the word has qu
	if strings.Contains(word, "qu") || strings.HasPrefix(word, "qu") {
		index := strings.LastIndex(word, "qu") + 2
		return word[index:] + word[:index] + "ay"
	}
	return ""
}

func Rule4(word string) string {
	// if the word starts with consonant and has y
	for _, char := range word {
		if !strings.ContainsRune(vowels, char) {
			if strings.Contains(word, "y") || strings.HasPrefix(word, "y") {
				index := strings.LastIndex(word, "y")
				return word[index:] + word[:index] + "ay"
			}
		}
	}
	return ""
}
