package yacht

import (
	"sort"
)

// Score calculates the score for a given dice roll and category
func Score(dice []int, category string) int {
	switch category {
	case "ones":
		return sumFace(dice, 1)
	case "twos":
		return sumFace(dice, 2)
	case "threes":
		return sumFace(dice, 3)
	case "fours":
		return sumFace(dice, 4)
	case "fives":
		return sumFace(dice, 5)
	case "sixes":
		return sumFace(dice, 6)

	case "yacht":
		if allEqual(dice) {
			return 50
		}
		return 0

	case "full house":
		if isFullHouse(dice) {
			return sumAll(dice)
		}
		return 0

	case "four of a kind":
		if val, ok := fourOfAKind(dice); ok {
			return val * 4
		}
		return 0

	case "little straight":
		if isStraight(dice, []int{1, 2, 3, 4, 5}) {
			return 30
		}
		return 0

	case "big straight":
		if isStraight(dice, []int{2, 3, 4, 5, 6}) {
			return 30
		}
		return 0

	case "choice":
		return sumAll(dice)

	default:
		return 0
	}
}

// --- Helper functions ---

func sumFace(dice []int, face int) int {
	sum := 0
	for _, d := range dice {
		if d == face {
			sum += face
		}
	}
	return sum
}

func sumAll(dice []int) int {
	sum := 0
	for _, d := range dice {
		sum += d
	}
	return sum
}

func allEqual(dice []int) bool {
	for _, d := range dice {
		if d != dice[0] {
			return false
		}
	}
	return true
}

func isFullHouse(dice []int) bool {
	counts := make(map[int]int)
	for _, d := range dice {
		counts[d]++
	}
	if len(counts) == 2 {
		hasTwo, hasThree := false, false
		for _, c := range counts {
			if c == 2 {
				hasTwo = true
			}
			if c == 3 {
				hasThree = true
			}
		}
		return hasTwo && hasThree
	}
	return false
}

func fourOfAKind(dice []int) (int, bool) {
	counts := make(map[int]int)
	for _, d := range dice {
		counts[d]++
	}
	for face, c := range counts {
		if c >= 4 {
			return face, true
		}
	}
	return 0, false
}

func isStraight(dice []int, target []int) bool {
	sort.Ints(dice)
	for i := 0; i < len(dice); i++ {
		if dice[i] != target[i] {
			return false
		}
	}
	return true
}
