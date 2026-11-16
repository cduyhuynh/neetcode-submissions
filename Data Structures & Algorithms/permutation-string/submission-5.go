func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1_map := make(map[byte]int, 26)
	for i:=0; i<len(s1); i++{
		s1_map[s1[i]]++
	}

	s2_map := make(map[byte]int, 26)
	window := len(s1) - 1
	for i:=0; i<len(s2);i++{
		char := s2[i]
		s2_map[char]++
		if i < window {
			continue
		}
		contained := true
		for k, f := range(s1_map) {
			if s2_map[k] != f {
				contained = false
			}
		}
		if contained {
			return true
		}
		s2_map[s2[i-window]]--
	}

	return false
}
