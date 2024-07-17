func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	if target < matrix[0][0] || matrix[m-1][n-1] < target {
		return false
	}

	r := 0
	for r < m && matrix[r][0] <= target {
		r++
	}
	r--

	i, j := 0, n-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case matrix[r][med] < target:
			i = med + 1
		case target < matrix[r][med]:
			j = med - 1
		default:
			return true
		}
	}

	return matrix[r][j] == target
}