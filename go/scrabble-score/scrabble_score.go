package scrabble

import "strings"

func Score(word string) int {
	// panic("Please implement the Score function")
	scrableMap := map[string]int{
		// A, E, I, O, U, L, N, R, S, T
		"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1,
		// D, G
		"D": 2, "G": 2,
		//   B, C, M, P
		"B": 3, "C": 3,"M": 3,"P": 3, 
		// F, H, V, W, Y
		"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
		// K
		"K": 5,
		// J, X
		"J": 8, "X": 8,
		// Q, Z
		"Q": 10, "Z": 10,
	}

	score := 0
	for _, letter := range strings.ToUpper(word) {
		if val, cond := scrableMap[string(letter)]; cond == true {
			score += val
		}
	}
	return score
}
