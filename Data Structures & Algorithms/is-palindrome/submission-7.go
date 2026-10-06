func isPalindrome(s string) bool {
	var alphanumeric_runes []rune
	for _, char := range(s) {
		if unicode.IsLetter(char) {
			alphanumeric_runes = append(alphanumeric_runes, unicode.ToLower(char))
		} else if unicode.IsDigit(char) {
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
