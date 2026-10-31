func isValid(s string) bool {
	s_stack := []byte{s[0]}
	for i:=1; i<len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			s_stack = append(s_stack, c)
		} else {
			if len(s_stack) == 0 {
				return false
			}
			tmp := s_stack[len(s_stack)-1]
			if (c == ')' && tmp == '(') || 
				(c == '}' && tmp == '{') || 
				(c == ']' && tmp == '[') {
				s_stack = s_stack[:len(s_stack)-1]
			} else {
				return false
			}
		}
	}
	return len(s_stack) == 0
}
