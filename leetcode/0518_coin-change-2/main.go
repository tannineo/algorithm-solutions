package main

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func change(amount int, coins []int) int {

	lenCoin := len(coins)

	dp := make([]int, amount+1)

	dp[0] = 1

	sort.Ints(coins)

	for i := 0; i < lenCoin; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}

	return dp[amount]
}

func main() {
	println(change(5, []int{1, 2, 5})) // 4

	println(change(3, []int{2})) // 0

	println(change(10, []int{10})) // 1

}
