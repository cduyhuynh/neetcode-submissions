func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int)
	for i := 0; i < len(nums); i++{
		to_be_found := target - nums[i]
		if idx, found := nmap[to_be_found]; found {
			return []int{idx, i}
		} else {
			nmap[nums[i]] = i
		}
	}
	return []int{}
}
