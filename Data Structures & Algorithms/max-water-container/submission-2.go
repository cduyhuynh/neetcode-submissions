func maxArea(heights []int) int {
	max_area := 0
	l,r := 0,len(heights) - 1
	for l < r {
		max_area = max(max_area, (min(heights[l], heights[r]))*(r - l))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return max_area
}