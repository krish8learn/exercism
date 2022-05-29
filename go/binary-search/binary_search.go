package binarysearch

func SearchInts(list []int, key int) int {
	// panic("Please implement the SearchInts function")
	left := 0
	right := len(list) - 1
	// will catch the index on which element found
	retIndex := -1
	for left <= right {
		//middle of slice
		middle := left + (right-left)/2
		if list[middle] == key {
			//key found
			retIndex = middle
			break
		} else if list[middle] > key {
			// key is less than middle element, find middle of lower half
			right = middle - 1
		} else if list[middle] < key {
			// key is higher than element, find middle of higher half
			left = middle + 1
		}
	}

	return retIndex

}
