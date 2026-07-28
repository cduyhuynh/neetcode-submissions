func isAnagram(s string, t string) bool {
	var slen = len(s)
	var tlen = len(t)
	if (slen != tlen) {
		return false
	}
	smap := make(map[rune]int, 26)
	for _, sval := range(s) {
		smap[sval] += 1
	}

	for _, tval := range(t) {
		if smap[tval] == 0 {
			return false
		} else {
			smap[tval] -= 1
		}
	}

	return true
}
