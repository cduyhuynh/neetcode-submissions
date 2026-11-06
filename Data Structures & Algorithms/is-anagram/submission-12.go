func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s_map := make(map[byte]int)

	for i:=0; i<len(s); i++{
		s_map[s[i]]++
	}
	for i:=0; i<len(t); i++{
		s_map[t[i]]--
		if s_map[t[i]] < 0 { 
			return false 
		}
	}
	return true
}
