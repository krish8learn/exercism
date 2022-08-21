package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

var list []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	// panic("Please implement the Foldl function")
	if s.Length() == 0 {
		return initial
	}
	return s[1:].Foldl(fn, fn(initial, s[0]))
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	// panic("Please implement the Foldr function")
	if s.Length() == 0 {
		return initial
	}
	return s[:s.Length()-1].Foldr(fn, fn(s[s.Length()-1],initial))
}

func (s IntList) Filter(fn func(int) bool) IntList {
	// panic("Please implement the Filter function")
	result := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (s IntList) Length() int {
	// panic("Please implement the Length function")
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	// panic("Please implement the Map function")

	if s.Length() == 0 {
		return s
	}

	result := make(IntList, 0)
	for _, value := range s {
		result = append(result, fn(value))
	}
	return result
}

func (s IntList) Reverse() IntList {
	// panic("Please implement the Reverse function")

	//check empty list
	sLength := len(s)
	if sLength == 0 {
		return s
	}

	for i := 0; i < sLength/2; i++ {
		j := sLength - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	// panic("Please implement the Append function")
	s = append(s, lst...)
	return s
}

func (s IntList) Concat(lists []IntList) IntList {
	// panic("Please implement the Concat function")

	for _, innerList := range lists {
		for _, element := range innerList {
			s = append(s, element)
		}
	}

	return s
}
