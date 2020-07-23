package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	len1 := len(text1)
	len2 := len(text2)

	if text1 == text2 {
		return len1
	}

	// init dp
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for j := range dp[0] {
		dp[0][j] = 0
	}
	for i := range dp {
		dp[i][0] = 0
	}

	// calculate dp
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len1][len2]
}

func main() {
	println(longestCommonSubsequence("abcde", "ace")) // 3

	println(longestCommonSubsequence("abc", "abc")) // 3

	println(longestCommonSubsequence("abc", "def")) // 0

}
