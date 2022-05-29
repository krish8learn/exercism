package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	// panic("Implement the Distance function")

	input1Len, input2Len := len(a), len(b)

	if input1Len != input2Len {
		//inputs lengths are not same, error
		return 0, errors.New("input lengths are different")
	} else if input1Len == 0 || input2Len == 0 || strings.EqualFold(a, b) {
		//any one of inputs are not empty or same
		return 0, nil
	}

	//inputs lengths are same and different
	count := 0
	for index := 0; index < input1Len; index++ {
		if a[index] != b[index] {
			count++
		}
	}
	return count, nil
}
