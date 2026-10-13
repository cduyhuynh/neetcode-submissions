func maxProfit(prices []int) int {
	max_profits := 0
	buy := prices[0]
	for i:=1; i<len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			continue
		}
		max_profits = max(max_profits, prices[i] - buy)
	}
	return max_profits
}
