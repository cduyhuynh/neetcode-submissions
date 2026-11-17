// 2,9,8,3,6
// root
// 2 9
// 8 3
// 6
func rob(nums []int) int {
	lnums := len(nums)
	memo := make([]int, lnums+1)
	for i:=0; i<=lnums; i++{
		memo[i] = -1
	}

    var dfs func(int) int
	dfs = func(idx int) int{
		if idx >= lnums {
			return 0
		}
		if memo[idx] != -1 {
			return memo[idx]
		}
		memo[idx] = max(nums[idx] + dfs(idx+2), dfs(idx+1))
		return memo[idx]
	}
	return dfs(0)
}
