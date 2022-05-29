package lsproduct

import (
	"errors"

	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	// panic("Please implement the LargestSeriesProduct function")

	//input error checking
	inputLen := len(digits)

	if span < 0 {
		return -1, errors.New("span must not be negative")
	} else if span > inputLen {
		return -1, errors.New("span must be smaller than string length")
	} else if err := numberValidtyCheck(digits); err != nil {
		// log.Println(err)
		return -1, errors.New("digits input must only contain digits")
	}

	//possible permutation
	perm := inputLen - span + 1
	var max_product int64

	for i := 0; i < perm; i++ {
		product := calculateProduct(digits[i : i+span])
		if product > max_product {
			max_product = product
		}
	}

	return max_product, nil
}

func calculateProduct(subString string) int64 {
	product := 1
	for _, value := range subString {
		number, _ := strconv.Atoi(string(value))
		product = product * number
	}
	return int64(product)
}

func numberValidtyCheck(digits string) error {

	for _, value := range digits {
		if _, err := strconv.Atoi(string(value)); err != nil {
			return errors.New("digits input must only contain digits")
		}
	}
	return nil
}
