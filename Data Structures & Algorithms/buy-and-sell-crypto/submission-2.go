func maxProfit(prices []int) int {
	max_profit := 0
	for i:=0; i<len(prices)-1; i++ {
		for j:=i+1; j<len(prices); j++{
			tmp := prices[j] - prices[i]
			max_profit = max(max_profit, tmp)
		}
	}
	return max_profit
}
