func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	if target < nums[l] || target > nums[r] {
		return -1
	}

	for l <= r {
		m := (l+r)/2
		fmt.Println("m", m, "nums[m]", nums[m])
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
		fmt.Println("l", l, "r", r)
	}

	return -1
}
