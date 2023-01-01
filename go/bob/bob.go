// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	remark = strings.TrimSpace(remark)
	if remark == "" || onlySpecialCharacters(remark) && remark[len(remark)-1] != '?' {
		// for every character is special characters -> "Fine. Be that way!"
		return "Fine. Be that way!"
	} else if remark[len(remark)-1] == '?' && remark == strings.ToUpper(remark) && noDigit(remark) && mustContainLetter(remark) {
		// for yelling a question
		return "Calm down, I know what I'm doing!"
	} else if remark[len(remark)-1] == '?' {
		// sentence ends with question ?-> "Sure"
		return "Sure."
	} else if !onlyDigit(remark) && remark == strings.ToUpper(remark) {
		// shout(everything in capital) -> "Whoa, chill out!"
		return "Whoa, chill out!"
	}
	// anything else -> "Whatever"
	return "Whatever."
}

func isUpperAll(input string) bool {
	for _, char := range input {
		if unicode.IsLower(char) {
			return false
		}
	}
	return true
}

func onlySpecialCharacters(input string) bool {
	//every letter is special character (no letter and digit)
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func onlyDigit(input string) bool {
	//every character is digit
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ",", "")
	for _, char := range input {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func noDigit(input string) bool {
	for _, char := range input {
		if unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func mustContainLetter(input string) bool {
	for _, char := range input {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}
