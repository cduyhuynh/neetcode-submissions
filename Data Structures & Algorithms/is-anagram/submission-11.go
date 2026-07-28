func isAnagram(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}
	smap := make(map[byte]int, 26)
	for i := 0; i < len(s); i++{
		smap[s[i]]++
	}

	for i := 0; i < len(t); i++{
		if smap[t[i]] == 0 {
			return false
		} else {
			smap[t[i]]--
		}
	}

	return true
}
