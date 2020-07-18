package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}

	dp[0][0] = 0
	for i, v := range s1 {
		dp[i+1][0] = dp[i][0] + int(v)
	}
	for j, v := range s2 {
		dp[0][j+1] = dp[0][j] + int(v)
	}

	for i, v1 := range s1 {
		for j, v2 := range s2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1]+int(v1), dp[i+1][j]+int(v2))
			}
		}
	}

	return dp[len(s1)][len(s2)]
}

func main() {
	println(minimumDeleteSum("sea", "eat")) // 231

	println(minimumDeleteSum("delete", "leet")) // 403

}
