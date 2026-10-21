func numIslands(grid [][]byte) int {
    num := 0
	m := len(grid)
	n := len(grid[0])
	directions := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	var traversal func(int, int)
	traversal = func(i int, j int) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		cell := grid[i][j]
		if cell == '1' {
			grid[i][j] = '0'
			for _, direction := range(directions) {
				traversal(i + direction[0], j + direction[1])
			}
		}
	}

	for i:=0; i<m;i++{
		for j:=0; j<n; j++{
			cell := grid[i][j]
			if cell == '1' {
				num++
				traversal(i, j)
			}
		}
	}
	return num
}
