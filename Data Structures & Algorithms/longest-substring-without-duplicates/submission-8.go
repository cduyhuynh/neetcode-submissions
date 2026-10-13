func lengthOfLongestSubstring(s string) int {
	length := 0
	hash := make(map[byte]int)
	for l,r := 0,0; r < len(s); r++{
		letter := s[r]
		hash[letter]++
		if hash[letter] == 1 {
			length = max(length, r - l + 1)
			continue
		} else {
			length = max(length, r - l - 1)
			for l < r {
				hash[s[l]]--
				if s[l] == letter {
					l++
					break
				} else {
					l++
				}
			}
		}
	}
	return length
}
