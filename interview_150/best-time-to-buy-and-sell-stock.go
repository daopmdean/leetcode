package interview_150

func maxProfit(prices []int) int {
	buy := prices[0]
	max := 0

	for _, price := range prices {
		if price < buy {
			buy = price
		} else if price-buy > max {
			max = price - buy
		}
	}

	return max
}
