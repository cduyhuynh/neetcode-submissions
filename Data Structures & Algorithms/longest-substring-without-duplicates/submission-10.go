func lengthOfLongestSubstring(s string) int {
	length := 0
	hash := make(map[byte]int)
	for l,r := 0,0; r < len(s); r++{
		if idx, found := hash[s[r]]; found {
			l = max(l, idx + 1)
		}
		hash[s[r]] = r
		length = max(length, r - l + 1)
	}
	return length
}
