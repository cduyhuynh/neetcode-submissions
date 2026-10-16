func characterReplacement(s string, k int) int {
	smap := make(map[byte]int)
	longest, l, maxf := 0, 0, 0
	for r := 0; r < len(s); r++ {
		smap[s[r]]++
		if smap[s[r]] > maxf {
			maxf = smap[s[r]]
		}
		for r - l + 1 - maxf > k {
			smap[s[l]]--
			l++
		}
		longest = max(longest, r - l +1)
	}
	return longest
}

