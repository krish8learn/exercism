package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	// panic("Please implement the Application() function")
	recommendation, _ := utf8.DecodeRuneInString("❗")
	search, _ := utf8.DecodeRuneInString("🔍")
	weather, _ := utf8.DecodeRuneInString("☀")
	// found := false
	for _, value := range log {
		if value == recommendation {
			return "recommendation"
		} else if value == search {
			return "search"
		} else if value == weather {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// panic("Please implement the Replace() function")
	runeSlice := []rune{}
	for _, value := range log {
		if value == oldRune {
			runeSlice = append(runeSlice, newRune)
		} else {
			runeSlice = append(runeSlice, value)
		}
	}
	return string(runeSlice)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// panic("Please implement the WithinLimit() function")
	recommendation, _ := utf8.DecodeRuneInString("❗")
	for index, value := range log {
		if value == recommendation && index < limit {
			return true
		}
	}
	return false
}
