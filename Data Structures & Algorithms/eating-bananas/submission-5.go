func minEatingSpeed(piles []int, h int) int {
	speed := 1
	min_speed := 1
	max_speed := 1
	for _, pile := range(piles) {
		if pile > max_speed {
			max_speed = pile
		}
	}
	fmt.Println("max_speed", max_speed)

	for min_speed <= max_speed {
		hour := 0
		mid := (min_speed+max_speed)/2
		for _, pile := range(piles) {
			if pile % mid == 0 {
				hour += pile/mid	
			} else {
				hour += pile/mid + 1
			}
		}
		fmt.Println("speed", speed, "hour", hour)
		if hour > h {
			min_speed = mid + 1
		} else {
			max_speed = mid - 1
			speed = mid
		}
	}
	return speed
}
