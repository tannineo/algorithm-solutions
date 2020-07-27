package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxUncrossedLines(A []int, B []int) int {
	lenA := len(A)
	lenB := len(B)

	// init dp
	dp := make([][]int, lenA+1)
	for i := 0; i <= lenA; i++ {
		dp[i] = make([]int, lenB+1)
		dp[i][0] = 0
	}
	for j := 0; j <= lenB; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[lenA][lenB]
}

func main() {
	println(maxUncrossedLines([]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2})) // 3
}
