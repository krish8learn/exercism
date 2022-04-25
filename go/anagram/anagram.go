package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	// panic("Please implement the Detect function")

	var returnStringSlice []string

	for _, valueCandidates := range candidates {
		if len(valueCandidates) == len(subject) {
			if !strings.EqualFold(subject, valueCandidates) {
				subjectFrequency := LetterFrequency(strings.ToLower(subject))
				candidateFrequency := LetterFrequency(strings.ToLower(valueCandidates))
				if reflect.DeepEqual(subjectFrequency, candidateFrequency) {
					returnStringSlice = append(returnStringSlice, valueCandidates)
				}
			}

		}
	}
	return returnStringSlice
}

func LetterFrequency(word string) map[string]int {

	frequency := make(map[string]int)

	for _, value := range word {
		frequency[string(value)]++
	}
	return frequency
}
