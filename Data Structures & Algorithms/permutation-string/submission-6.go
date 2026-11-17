func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1_arr := make([]int, 26)
	s2_arr := make([]int, 26)
	for i:=0; i<len(s1);i++{
		s1_arr[s1[i] - 'a']++
		s2_arr[s2[i] - 'a']++
	}

	matches := 0
	for i:=0; i<26; i++{
		if s1_arr[i] == s2_arr[i] {
			matches++
		}
	}

	for l,r:=0,len(s1); r<len(s2); r++{
		if matches == 26 {
			return true
		}

		char := s2[r] - 'a'
		s2_arr[char]++
		if s1_arr[char] == s2_arr[char] {
			matches++
		} else if s1_arr[char]+1 == s2_arr[char] {
			matches--
		}

		char = s2[l] - 'a'
		s2_arr[char]--
		if s1_arr[char] == s2_arr[char]{
			matches++
		} else if s1_arr[char]-1 == s2_arr[char] {
			matches--
		}
		l++
	}

	return matches == 26
}
