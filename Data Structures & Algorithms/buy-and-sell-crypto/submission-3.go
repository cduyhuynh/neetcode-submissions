import (
	"slices"
)
func maxProfit(prices []int) int {
	var profits []int
	buy := prices[0]
	for i:=1; i<len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			continue
		}
		profits = append(profits, prices[i] - buy)
	}
	if len(profits) == 0 {
		return 0
	} else {
		return slices.Max(profits)
	}
}
