package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	// panic("Please implement the Transform function")
	ret := make(map[string]int, 0)

	for intValue, slice := range in {
		for _, stringValues := range slice {
			ret[strings.ToLower(stringValues)] = intValue
		}
	}
	return ret
}
