func search(nums []int, target int) int {
	l := 0
	r := len(nums)-1
	for l <= r {
		idx := (l + r)/2
		if nums[idx] > target {
			r = idx - 1
		} else if nums[idx] < target{
			l = idx + 1
		} else {
			return idx
		}
	}
	return -1
}
