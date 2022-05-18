package prime

func Factors(n int64) []int64 {
	// panic("Please implement the Factors function")
	divisors := make([]int64, 0)
	getPrimes(n, &divisors, 2)
	return divisors

}

func getPrimes(n int64, divisors *[]int64, div int64) {
	if n == 1 {
		return
	}

	if n%div == 0 {
		*divisors = append(*divisors, div)
		getPrimes(n/div, divisors, div)
	} else {
		getPrimes(n, divisors, div+1)
	}
}
