func trap(heights []int) int {
	area := 0
	n := len(heights)
	prefix := make([]int, n)
	suffix := make([]int, n)

	prefix[0] = heights[0]
	for i:=1; i < n; i++ {
		if heights[i] > prefix[i-1] {
			prefix[i] = heights[i]
		} else {
			prefix[i] = prefix[i - 1]
		}
	}
	
	suffix[n-1] = heights[n-1]
	for i:=n-2; i >= 0; i-- {
		if heights[i] > suffix[i + 1] {
			suffix[i] = heights[i]
		} else {
			suffix[i] = suffix[i + 1]
		}
	}

	for i:=0; i < n; i++ {
		area += min(prefix[i], suffix[i]) - heights[i]
	}

	return area
}
