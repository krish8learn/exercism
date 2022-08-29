package sieve

func Sieve(limit int) []int {
	// panic("Please implement the Sieve function")

	var ret []int
	//create an array upto the limit
	number := make([]int, limit+1)

	for i := 2; i <= limit; i++ {
		if number[i] == 0 {
			//not marked, must be prime
			ret = append(ret, i)
			for j := i; j <= limit; j++ {
				//marking all the multiples under the limit
				if j*i <= limit {
					//not prime, so marked
					number[j*i] = 1
				}
			}
		}
	}

	return ret
}
