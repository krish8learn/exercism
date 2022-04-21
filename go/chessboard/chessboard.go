package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	// panic("Please implement CountInRank()")
	count := 0
	for _, value := range cb[rank] {
		if value == true {
			count++
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	// panic("Please implement CountInFile()")
	count := 0
	if file >= 1 && file <= 8 {
		for _, value1 := range cb {
			if value1[file-1] == true {
				count++
			}
		}
		return count
	}
	return 0
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	// panic("Please implement CountAll()")
	return 8 * 8
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	// panic("Please implement CountOccupied()")
	count := 0
	for _, value1 := range cb {
		for _, value2 := range value1 {
			if value2 == true {
				count++
			}
		}
	}
	return count
}
