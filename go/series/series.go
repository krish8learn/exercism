package series

func All(n int, s string) []string {
	// panic("Please implement the All function")
	result := make([]string, 0, 0)
	stringLength := len(s)

	for index := range s {
		if index+n <= stringLength {
			result = append(result, s[index:index+n])
		}
	}

	return result
}

func UnsafeFirst(n int, s string) string {
	// panic("Please implement the UnsafeFirst function")
	output := ""
	if n <= len(s) {
		output = s[0:n]
	}
	return output
}
