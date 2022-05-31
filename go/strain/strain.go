package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	// panic("Please implement the Keep function")
	var keepInts Ints
	for _, value := range i {
		if filter(value) {
			keepInts = append(keepInts, value)
		}
	}
	return keepInts
}

func (i Ints) Discard(filter func(int) bool) Ints {
	// panic("Please implement the Discard function")
	var discardInts Ints
	for _, value := range i {
		if !filter(value) {
			discardInts = append(discardInts, value)
		}
	}
	return discardInts
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	// panic("Please implement the Keep function")
	var listKeep Lists
	for _, value := range l {
		if filter(value) {
			listKeep = append(listKeep, value)
		}
	}
	return listKeep

}

func (s Strings) Keep(filter func(string) bool) Strings {
	// panic("Please implement the Keep function")
	var stringKeep Strings
	for _, value := range s {
		if filter(value) {
			stringKeep = append(stringKeep, value)
		}
	}
	return stringKeep
}
