// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	//remove special characters except  , ',: from string and replace with space//
	re, err := regexp.Compile(`[^a-zA-Z\d\s:']`)
	if err != nil {
		log.Fatal(err)
	}
	s = re.ReplaceAllString(s, " ")
	//split the string based on space//
	words := strings.Split(s, " ")
	var acronym string
	for _, word := range words {
		if word == "" {
			continue
		}
		acronym += strings.ToUpper(string(word[0]))
	}
	return acronym
}
