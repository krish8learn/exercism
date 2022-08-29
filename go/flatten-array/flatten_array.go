package flatten

func Flatten(nested interface{}) []interface{} {
	// panic("Please implement the Flatten function")
	result := []interface{}{}

	switch insideItem := nested.(type) {
	case []interface{}:
		for _, value := range insideItem {
			result = append(result, Flatten(value)...)
		}
	case interface{}:
		result = append(result, insideItem)

	}
	return result
}
