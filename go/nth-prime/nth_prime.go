package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	// panic("Please implement the Nth function")
	if n < 1 {
		return 0, errors.New("input must be greater than 1")
	}

	num := 0
	for x := 2; x <= 2147483648; x++ {
		if isPrime(x) {
			num++
			if num == n {
				return x, nil
			}
		}
	}

	return 0, nil

}

func isPrime(x int) bool {

	if x == 2 {
		return true
	}

	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true

}
