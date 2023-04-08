package wordsearch

import "errors"

func Solve(words []string, puzzle []string) (results map[string][2][2]int, err error) {

	results = map[string][2][2]int{}

	for _, word := range words {

		wordLength := len(word)

		for r, line := range puzzle {

			for c := range line {

				if word[0:1] == line[c:c+1] {

					if slice := sliceRight(line, c, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c + wordLength - 1, r}}

						break

					}

					if slice := sliceLeft(line, c, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c - wordLength + 1, r}}

						break

					}

					if slice := sliceDown(puzzle, c, r, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c, r + wordLength - 1}}

						break

					}

					if slice := sliceUp(puzzle, c, r, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c, r - wordLength + 1}}

						break

					}

					if slice := sliceDownRight(puzzle, c, r, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c + wordLength - 1, r + wordLength - 1}}

						break

					}

					if slice := sliceDownLeft(puzzle, c, r, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c - wordLength + 1, r + wordLength - 1}}

						break

					}

					if slice := sliceUpRight(puzzle, c, r, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c + wordLength - 1, r - wordLength + 1}}

						break

					}

					if slice := sliceUpLeft(puzzle, c, r, wordLength); slice == word {

						results[word] = [2][2]int{{c, r}, {c - wordLength + 1, r - wordLength + 1}}

						break

					}

				}

			}

		}

	}

	if len(results) != len(words) {

		err = errors.New("word(s) not found")

	}

	return

}

func sliceDown(list []string, startX, startY, length int) string {

	height := len(list)

	if startY+length > height {

		return ""

	}

	var slice []byte

	for y := startY; y < len(list); y++ {

		slice = append(slice, list[y][startX])

	}

	return string(slice)

}

func sliceDownLeft(list []string, startX, startY, length int) string {

	height := len(list)

	if startX-length < 0 || startY+length > height {

		return ""

	}

	var slice []byte

	for x, y := startX, startY; x > 0 && x < startX+length && y < height-1; x, y = x-1, y+1 {

		slice = append(slice, list[y][x])

	}

	return string(slice)

}

func sliceDownRight(list []string, startX, startY, length int) string {

	height := len(list)

	width := len(list[0])

	if startX+length > width || startY+length > height {

		return ""

	}

	var slice []byte

	for x, y := startX, startY; x < startX+length && y < height; x, y = x+1, y+1 {

		slice = append(slice, list[y][x])

	}

	return string(slice)

}

func sliceLeft(line string, start int, length int) string {

	if start-length+1 < 0 {

		return ""

	}

	var slice []byte

	for x := start; x > start-length; x-- {

		slice = append(slice, line[x])

	}

	return string(slice)

}

func sliceRight(line string, start int, length int) string {

	if len(line)-start < length {

		return ""

	}

	return string(line[start : start+length])

}

func sliceUp(list []string, startX, startY, length int) string {

	if startY-length+1 < 0 {

		return ""

	}

	var slice []byte

	for y := startY; y > startY-length; y-- {

		slice = append(slice, list[y][startX])

	}

	return string(slice)

}

func sliceUpLeft(list []string, startX, startY, length int) string {

	if startX-length+1 < 0 || startY-length+1 < 0 {

		return ""

	}

	var slice []byte

	for x, y := startX, startY; x > 0 && x > startX-length && y > 0 && y > startY-length; x, y = x-1, y-1 {

		slice = append(slice, list[y][x])

	}

	return string(slice)

}

func sliceUpRight(list []string, startX, startY, length int) string {

	width := len(list[0])

	if startX+length > width || startY-length+1 < 0 {

		return ""

	}

	var slice []byte

	for x, y := startX, startY; x < startX+length && y > 0; x, y = x+1, y-1 {

		slice = append(slice, list[y][x])

	}

	return string(slice)

}
