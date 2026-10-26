func climbStairs(n int) int {
	steps := make([]int, n+1)
	for i := 0; i <= n; i++{
		steps[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			if i == n {
				return 1
			}
			return 0
		}
		if steps[i] != -1 {
			return steps[i]
		}
		steps[i] = dfs(i+1) + dfs(i+2)
		return steps[i]
	}
	return dfs(0)
}

