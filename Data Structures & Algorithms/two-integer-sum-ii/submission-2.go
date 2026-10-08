func twoSum(numbers []int, target int) []int {
	idx1 := 0
	idx2 := len(numbers) - 1
	for idx1 < idx2{
		if (numbers[idx1] + numbers[idx2]) == target {
			return []int{idx1 + 1, idx2 + 1}
		}
		if (numbers[idx1] + numbers[idx2]) > target {
			idx2--
		} else {
			idx1++
		}
	}
	return []int{0,0}
}
