func topKFrequent(nums []int, k int) []int {
	var k_frqs []int
	nums_map := make(map[int]int)
	for _, num := range(nums) {
		nums_map[num]++
	}

	lnums := len(nums)
	nums_bucket := make([][]int, lnums + 1)
	for num, frq := range(nums_map) {
		nums_bucket[frq] = append(nums_bucket[frq], num)
	}

	for i:=lnums; i >= 0 && k > 0; i-- {
		k_frqs = append(k_frqs, nums_bucket[i]...)
		k -= len(nums_bucket[i])
	}

	return k_frqs
}
