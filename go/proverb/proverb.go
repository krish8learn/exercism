// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "strings"

// Proverb should have a comment documenting it.
var lines = []string{
	"And all for the want of a _.",
	"For want of a _ the _ was lost.",
}

func Proverb(rhyme []string) []string {

	var output []string
	var length = len(rhyme)

	//check input length
	if length == 0 {
		return []string{}
	} else if length == 1 {
		got := strings.Replace(lines[0], "_", rhyme[0], -1)
		output = append(output, got)
		return output
	}

	for LineIndex, line := range lines {
		for rhymeIndex := range rhyme {
			if LineIndex == 0 {
				got := strings.Replace(line, "_", rhyme[rhymeIndex], -1)
				output = append(output, got)
				break
			}
			if rhymeIndex < length-1 {
				got := strings.Replace(line, "_", rhyme[rhymeIndex], 1)
				got = strings.Replace(got, "_", rhyme[rhymeIndex+1], 2)
				output = append(output, got)
			}
		}

	}

	return ChangeFirstElement(output)
}

func ChangeFirstElement(slice []string) []string {
	// var changed []string
	changed := slice[1:]
	changed = append(changed, slice[0])
	return changed
}
