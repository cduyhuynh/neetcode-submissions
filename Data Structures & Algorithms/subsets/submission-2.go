func subsets(nums []int) [][]int {
	lnums := len(nums)
	var subsets [][]int
	var subset []int
	var traversal func(int)
	traversal = func(idx int) {
		if idx >= lnums {
			tmp := make([]int, len(subset))
			copy(tmp, subset)
			subsets = append(subsets, tmp)
			return
		}
		subset = append(subset, nums[idx])
		traversal(idx+1)
		subset = subset[:len(subset) - 1]
		traversal(idx+1)
	}
	traversal(0)
	return subsets
}
