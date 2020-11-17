package main

func minCostClimbingStairs(cost []int) int {
	// cost elements number >= 2
	lastTwoStep1, lastTwoStep2 := cost[0], cost[1]

	cur := 2

	for cur < len(cost) {
		newStep := min(lastTwoStep1, lastTwoStep2) + cost[cur]
		lastTwoStep1, lastTwoStep2 = lastTwoStep2, newStep
		cur++
	}

	return min(lastTwoStep1, lastTwoStep2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
