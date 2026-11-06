func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	var anagrams [][]string

	charsMap := charsMap()
	for _, str := range(strs) {
		aMap := make(map[byte]int)
		key := ""
		for i:=0; i < len(str); i++{
			aMap[str[i]]++
		}
		for _, c := range(charsMap) {
			if aMap[c] > 0 {
				key += string(c) + string(aMap[c])
			}
		}
		groups[key] = append(groups[key], str)
	}
	for _, group := range(groups) {
		anagrams = append(anagrams, group)
	}
	
	return anagrams
}

func charsMap() []byte{
	var charMaps []byte
	var i byte
	for i = 'a'; i >= 'a' && i <= 'z'; i++ {
		charMaps = append(charMaps, i)
	}
	return charMaps
}
