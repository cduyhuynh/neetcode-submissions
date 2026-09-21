func topKFrequent(nums []int, k int) []int {
	num_map := make(map[int]int)
	for _, num := range(nums) {
		num_map[num]++
	}

	sorted_map := make(map[int][]int)
	for num, fre := range(num_map) {
		sorted_map[fre] = append(sorted_map[fre], num)
	}

	var k_array []int
	for i := len(nums); k > 0; i--{
		k_array = append(k_array, sorted_map[i]...)
		k -= len(sorted_map[i])
	}

	return k_array
}
