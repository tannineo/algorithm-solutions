package main

func maxProfit(prices []int) int {
	lenPrices := len(prices)
	if lenPrices <= 1 {
		return 0
	}

	profit := 0
	sellAt := len(prices) - 1

	for ; sellAt-1 >= 0; sellAt-- {
		// 1. find the peak value
		isPeak := false
		if sellAt == len(prices)-1 {
			isPeak = prices[sellAt-1] < prices[sellAt]
		} else {
			isPeak = prices[sellAt-1] < prices[sellAt] && prices[sellAt] >= prices[sellAt+1]
		}

		if isPeak {
			// 2. find the buyAt
			buyAt := sellAt - 1
			for ; buyAt-1 >= 0; buyAt-- {
				if prices[buyAt-1] > prices[buyAt] {
					break
				}
			}

			// 3. profit
			profit += prices[sellAt] - prices[buyAt]
			sellAt = buyAt
		}
	}

	return profit
}

func main() {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	println(maxProfit([]int{1, 2, 3, 4, 5}))
}
