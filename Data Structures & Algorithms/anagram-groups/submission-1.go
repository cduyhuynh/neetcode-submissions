import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	output := make(map[string][]string)
	for _, str := range strs {
		anagram := strConvert(str)
		output[anagram] = append(output[anagram], str)
	}
	var response [][]string
	for _, a := range output {
		response = append(response, a)
	}
	return response
}

func strConvert(str string) string {
	runes := []rune(str)
	slices.Sort(runes)
	tmp := string(runes)
	return tmp
}
