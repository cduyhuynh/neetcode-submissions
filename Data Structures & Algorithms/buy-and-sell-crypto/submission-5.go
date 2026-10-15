func maxProfit(prices []int) int {
	max_profit := 0
	buy := prices[0]
	for _, price := range(prices) {
		if price < buy {
			buy = price
			continue
		}
		max_profit = max(max_profit, price - buy)
	}
	return max_profit
}
