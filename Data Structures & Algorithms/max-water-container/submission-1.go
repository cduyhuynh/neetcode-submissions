func maxArea(heights []int) int {
	var max_water_amount int
	l,r := 0,len(heights) - 1
	for l < r {
		h := min(heights[l], heights[r])
		w := r - l
		max_water_amount = max(max_water_amount, h*w)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return max_water_amount
}