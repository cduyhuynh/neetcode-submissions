func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	tmp_sum := 0
	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			tmp_sum = nums[i] + nums[l] + nums[r]
			if tmp_sum > 0 {
				r--
			} else if tmp_sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1]{
					l++
				}
			}
		}
	}
	return res
}
