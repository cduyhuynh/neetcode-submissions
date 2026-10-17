func characterReplacement(s string, k int) int {
	longest := 0
	smap := make(map[byte]int)
	maxf := 0
	for l, r:=0,0; r < len(s); r++{
		smap[s[r]]++
		if smap[s[r]] > maxf {
			maxf = smap[s[r]]
		}
		if maxf + k < r - l + 1{
			smap[s[l]]--
			l++
		}
		if r - l + 1 > longest {
			longest = r - l + 1
		}
	}
	return longest
}

