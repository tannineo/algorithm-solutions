package main

import "math"

func numSquares(n int) int {
	// dp
	dp := make([]int, n+1)

	// init squares

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i-j*j]+1, dp[i])
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquaresMath(n int) int {
	// judge 1
	if isSquare(n) {
		return 1
	}

	// judge 4
	temp := n
	for temp%4 == 0 {
		temp /= 4
	}
	if temp%8 == 7 {
		return 4
	}

	// judge 2
	for i := 1; i*i < n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	// left 3
	return 3
}

func isSquare(n int) bool {
	temp := int(math.Sqrt(float64(n)))
	return temp*temp == n
}

func main() {
	println(numSquaresMath(15)) // 4

	println(numSquares(15)) // 4

}
