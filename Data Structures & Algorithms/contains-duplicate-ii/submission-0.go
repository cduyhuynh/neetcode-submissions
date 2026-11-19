func containsNearbyDuplicate(nums []int, k int) bool {
	nmap := make(map[int]int)
	for i:=0;i<len(nums);i++{
		if idx, found := nmap[nums[i]]; found {
			if i - idx <= k {
				return true
			}
		}
		nmap[nums[i]] = i
	}
	return false
}
