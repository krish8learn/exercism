package diamond

import (
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	// Check if character is A-Z
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("invalid input")
	}

	// Special case: single 'A'
	if char == 'A' {
		return "A", nil
	}

	// Calculate the size of the diamond
	size := int(char - 'A' + 1) // A=1, B=2, C=3, ..., Z=26
	width := 2*size - 1         // Width of the diamond

	var lines []string

	// Build the top half (including middle)
	for i := 0; i < size; i++ {
		letter := rune('A' + i)
		row := buildRow(letter, i, width)
		lines = append(lines, row)
	}

	// Build the bottom half (mirror of top, excluding middle)
	for i := size - 2; i >= 0; i-- {
		letter := rune('A' + i)
		row := buildRow(letter, i, width)
		lines = append(lines, row)
	}

	return strings.Join(lines, "\n"), nil
}

func buildRow(letter rune, index int, width int) string {
	if index == 0 {
		// First row: center the letter
		// For width=3: " A ", for width=5: "  A  ", etc.
		leftPad := (width - 1) / 2
		rightPad := width - leftPad - 1
		return strings.Repeat(" ", leftPad) + string(letter) + strings.Repeat(" ", rightPad)
	}

	// For other rows at index i (where i > 0):
	// Spaces before: i
	// First letter: 1
	// Spaces between: 2*i - 1
	// Second letter: 1
	// Spaces after: i
	// Total: i + 1 + (2*i - 1) + 1 + i = 4*i + 1... but width = 2*size - 1

	// The actual pattern should be:
	// Total used = 1 + (2*i - 1) + 1 = 2*i + 1
	// Padding on left: (width - (2*i + 1)) / 2 = size - i - 1
	// Padding on right: size - i - 1

	size := (width + 1) / 2
	padding := size - index - 1

	innerSpacing := 2*index - 1
	row := strings.Repeat(" ", padding) +
		string(letter) +
		strings.Repeat(" ", innerSpacing) +
		string(letter) +
		strings.Repeat(" ", padding)

	return row
}
