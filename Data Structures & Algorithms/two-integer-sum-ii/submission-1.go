func twoSum(numbers []int, target int) []int {
	idx1 := 0
	idx2 := len(numbers) - 1
	var arr []int
	for idx1 < idx2{
		if (numbers[idx1] + numbers[idx2]) == target {
			arr = append(arr, idx1 + 1, idx2 + 1)
			return arr
		}
		if (numbers[idx1] + numbers[idx2]) > target {
			idx2--
		} else {
			idx1++
		}
	}
	return arr
}
