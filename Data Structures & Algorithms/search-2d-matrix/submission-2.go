func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	l := 0
	r := m*n - 1
	for l <= r {
		mid := (l+r)/2
		tmp := matrix[mid/n][mid%n]
		fmt.Println("mid", mid, "tmp", tmp)
		if tmp > target {
			r = mid - 1
		} else if tmp < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}
