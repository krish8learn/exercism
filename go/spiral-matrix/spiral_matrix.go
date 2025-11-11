package spiralmatrix

func SpiralMatrix(size int) [][]int {

	// no input case
	if size == 0 {
		return [][]int{}
	}

	// create empty matrix
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	// define boundaries
	top, bottom := 0, size-1
	left, right := 0, size-1
	num := 1

	// fill the matrix in spiral order
	for top <= bottom && left <= right {
		// fill top row
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// fill right column
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// fill bottom row
		if top <= bottom {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = num
				num++
			}
			bottom--
		}

		// fill left column
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}
