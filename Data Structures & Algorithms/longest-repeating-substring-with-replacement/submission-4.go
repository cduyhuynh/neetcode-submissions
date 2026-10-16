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
		if r - l +1 > longest {
			longest = r - l +1
		}
	}
	return longest
}

