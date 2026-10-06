func isPalindrome(s string) bool {
	var alphanumeric_runes []rune
	for _, char := range(s) {
		if (char >= 'A' && char <= 'Z') {
			alphanumeric_runes = append(alphanumeric_runes, unicode.ToLower(char))
		}
		if ((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			alphanumeric_runes = append(alphanumeric_runes, char)
		}
	}
	fmt.Println(string(alphanumeric_runes))
	r := len(alphanumeric_runes) - 1
	for l := 0; l < len(alphanumeric_runes); l++{
		if alphanumeric_runes[l] != alphanumeric_runes[r] {
			return false
		} else {
			r--
		}
	}
	return true
}
