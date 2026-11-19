func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	longest := 1
	smap := make(map[byte]int)
	var maxf byte
	for l,r:=0,0; r<len(s);r++ {
		char := s[r]
		window := r-l+1
		smap[char]++
		if smap[char] > smap[maxf] {
			maxf = char
		}
		if smap[maxf] + k >= window {
			longest = max(longest, window)
			continue
		}
		smap[s[l]]--
		l++
	}
	return longest
}

