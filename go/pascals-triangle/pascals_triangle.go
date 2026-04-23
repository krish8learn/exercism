package pascal

func Triangle(n int) [][]int {
	triagnle := make([][]int, n)
	for i := 0; i < n; i++ {
		triagnle[i] = make([]int, i+1)
		triagnle[i][0], triagnle[i][i] = 1, 1 // first and last elements are always 1
		for j := 1; j < i; j++ {
			triagnle[i][j] = triagnle[i-1][j-1] + triagnle[i-1][j] // inner elements are sum of two above
		}
	}
	return triagnle
}
