func hasDuplicate(nums []int) bool {
    var maps = make(map[int]int)
		for _, num := range nums {
			if maps[num] == 1 {
				return true
			}
			maps[num] += 1
		}
		return false
}
