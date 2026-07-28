func isAnagram(s string, t string) bool {
	var slen = len(s)
	var tlen = len(t)
	if (len(s) != len(t)) {
		return false
	}
	smap := make(map[byte]int, 26)
	for i := 0; i < slen; i++{
		smap[s[i]]++
	}

	for i := 0; i < tlen; i++{
		if smap[t[i]] == 0 {
			return false
		} else {
			smap[t[i]]--
		}
	}

	return true
}
