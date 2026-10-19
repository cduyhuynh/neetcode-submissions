func minEatingSpeed(piles []int, h int) int {
	speed := 1
	min_speed := 1
	max_speed := 0
	for _, pile := range(piles) {
		max_speed = max(max_speed, pile)
	}
	for min_speed <= max_speed {
		time := 0
		mid := (min_speed + max_speed)/2
		for _, pile := range(piles) {
			if pile % mid > 0 {
				time += (pile/mid) + 1
			} else {
				time += pile/mid
			}
		}
		fmt.Println("time", time, "mid", mid)
		if time > h {
			min_speed = mid + 1
			speed = mid + 1
		} else if time <= h {
			max_speed = mid - 1
			speed = mid
		}
	}
	return speed
}
