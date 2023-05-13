package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	// panic("Please implement the CollatzConjecture function")

	count := 0
	if n <= 0 {
		return count, fmt.Errorf("error: number %v", n)
	}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			count++
		} else if n%2 == 1 {
			n = 3*n + 1
			count++
		}
	}

	return count, nil

}
