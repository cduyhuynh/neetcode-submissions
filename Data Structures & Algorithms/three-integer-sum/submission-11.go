func threeSum(nums []int) [][]int {
	var triplets [][]int
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums) - 1; i++ {
		l := i + 1
		r := len(nums) - 1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			tmp_sum := nums[i] + nums[l] + nums[r]
			if tmp_sum > 0 {
				r--
			} else if tmp_sum < 0 {
				l++
			} else {
				triplets = append(triplets, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return triplets
}
