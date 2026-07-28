func isAnagram(s string, t string) bool {
	var slen = len(s)
	var tlen = len(t)
	if (slen != tlen) {
		return false
	}
	smap := make(map[rune]int)
	tmap := make(map[rune]int)
	for _, sval := range(s) {
		smap[sval] += 1
	}

	for _, tval := range(t) {
		tmap[tval] += 1
	}
	for skey, sval := range(smap) {
		if sval != tmap[skey]{
			return false
		}
	}
	return true
}
