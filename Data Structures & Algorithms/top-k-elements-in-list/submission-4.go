func topKFrequent(nums []int, k int) []int {
	num_map := make(map[int]int)
	for _, num := range(nums) {
		num_map[num]++
	}

	sorted := make([][]int, len(nums) + 1)
	for num, fre := range(num_map) {
		sorted[fre] = append(sorted[fre], num)
	}

	var k_array []int
	for i := len(nums); k > 0; i--{
		k_array = append(k_array, sorted[i]...)
		k -= len(sorted[i])
	}

	return k_array
}
