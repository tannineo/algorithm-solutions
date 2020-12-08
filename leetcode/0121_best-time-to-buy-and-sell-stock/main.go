package main

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxPrice := 0
	maxProfit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		// update maxPrice
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		}

		// update maxProfit
		profit := maxPrice - prices[i]
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
