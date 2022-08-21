package sublist

type Relation string

func Sublist(l1, l2 []int) Relation {
	// panic("Please implement the Sublist function")
	switch {
	case len(l1) == len(l2) && isSub(l1, l2):
		return "equal"
	case len(l1) > len(l2) && isSub(l2, l1):
		return "superlist"
	case len(l1) < len(l2) && isSub(l1, l2):
		return "sublist"
	}
	return "unequal"
}

func isSub(firstInput, secondInput []int) bool {
	// return len(firstInput) == len(secondInput)

	for outer := 0; outer < len(secondInput)-len(firstInput)+1; outer++ {
		match := true
		for inner, value := range firstInput {
			if value != secondInput[inner+outer] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false

}
