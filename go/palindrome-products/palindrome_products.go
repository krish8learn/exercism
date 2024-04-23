package palindrome

import (
	"errors"
	"math"
)

// Define Product type here.
type Product struct {
	Product        int
	Factorizations [][2]int
}

func isPalindrome(n int) bool {
	orig := n
	reverseN := 0

	for {
		if n == 0 {
			break
		}
		reverseN = (reverseN * 10) + (n % 10)

		n /= 10
	}

	return orig == reverseN
}

func getFactors(n, min, max int) [][2]int {
	retVal := make([][2]int, 0)

	val := max

	for i := min; i < val; i++ {
		if n%i == 0 && n/i <= max {
			val = n / i
			retVal = append(retVal, [2]int{i, val})
		}
	}
	return retVal
}

func Products(fmin, fmax int) (Product, Product, error) {
	min := math.MaxInt
	max := 0

	minProduct := Product{}
	maxProduct := Product{}

	if fmin > fmax {
		return minProduct, maxProduct, errors.New("fmin > fmax...")
	}

	for i := fmin; i <= fmax; i++ {
		for j := fmin; j <= fmax; j++ {
			product := i * j

			if isPalindrome(product) {

				if product < min {
					min = product
				}

				if product > max {
					max = product
				}
			}
		}
	}

	if min != 1 {
		minProduct.Product = min
		minProduct.Factorizations = getFactors(min, fmin, fmax)
	}

	if max != 0 {
		maxProduct.Product = max
		maxProduct.Factorizations = getFactors(max, fmin, fmax)
	} else {
		return minProduct, maxProduct, errors.New("no palindromes...")
	}

	return minProduct, maxProduct, nil
}
