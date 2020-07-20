package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	lenM := len(grid)
	if lenM == 0 {
		return 0 // zero case?
	}
	lenN := len(grid[0])
	if lenN == 0 {
		return 0 // zero case?
	}

	dp := make([][]int, lenM)
	for i := range dp {
		dp[i] = make([]int, lenN)
	}

	// dp init value
	sum := 0
	for i := range grid {
		sum += grid[i][0]
		dp[i][0] = sum
	}
	sum = 0
	for i := range grid[0] {
		sum += grid[0][i]
		dp[0][i] = sum
	}

	// calculate dp
	for i := range dp {
		if i == 0 {
			continue
		}
		for j := range dp[i] {
			if j == 0 {
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[lenM-1][lenN-1]
}

func main() {
	println(minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	})) // 7
}
