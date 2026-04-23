package eliudseggs

import "strconv"

func EggCount(displayValue int) int {
	if displayValue == 0 {
		return 0
	}

	binaryStr := strconv.FormatInt(int64(displayValue), 2)
	eggs := 0
	for _, char := range binaryStr {
		if char == '1' {
			eggs++
		}
	}
	return eggs
}
