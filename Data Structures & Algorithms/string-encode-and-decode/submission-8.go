type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded string
	for _, str := range(strs) {
		encoded += strconv.Itoa(len(str))
		encoded += "#"
		encoded += str
	}
	// fmt.Println(encoded)
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	var decoded []string
	wcount := 0
	i:=0
	for i<len(encoded)-1 {
		j:=i+1
		for encoded[j] != '#'{
			j++
		}
		wcount, _ = strconv.Atoi(encoded[i:j])
		start := j+1
		end := min(start+wcount, len(encoded))
		// fmt.Println("start", start, "end", end)
		
		decoded = append(decoded, encoded[start:end])
		// fmt.Println("decoded", decoded)
		i = end
	}
	return decoded
}
