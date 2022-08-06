package pythagorean

import "math"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	//panic("Please implement the Range function")

	var triplets []Triplet
	//range over the min to max
	for first := min; first <= max; first++ {
		for second := first + 1; second <= max; second++ {
			if cond, value := squareRootSumCheckForRange(first, second, max); cond {
				triplets = append(triplets, Triplet{first, second, value})
			}
		}
	}

	return triplets
}

func squareRootSumCheckForRange(value1, value2, max int) (bool, int) {
	floatValue := math.Sqrt(float64(value1*value1 + value2*value2))

	if floatValue == float64(int(floatValue)) && int(floatValue) <= max {
		return true, int(floatValue)
	}

	return false, 0
}

func squareRootSumCheckForSum(value1, value2, max int) (bool, int) {
	floatValue := math.Sqrt(float64(value1*value1 + value2*value2))

	if floatValue == float64(int(floatValue)) && (value1 + value2 + int(floatValue)) == max {
		return true, int(floatValue)
	}

	return false, 0
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	//panic("Please implement the Sum function")

	var triplets []Triplet
	//range over the min to max
	for first := 1; first <= p; first++ {
		for second := first + 1; second <= p; second++ {
			if cond, value := squareRootSumCheckForSum(first, second, p); cond {
				triplets = append(triplets, Triplet{first, second, value})
			}
		}
	}

	return triplets

	
}
