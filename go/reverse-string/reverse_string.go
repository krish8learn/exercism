package reverse

func Reverse(input string) string {
	// panic("Please implement the Reverse function")
	result := ""
	for _, v := range input {
		result = string(v) + result
	}
	return result
}
