package main

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	len1 := len(word1)
	len2 := len(word2)
	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}

	// init dp
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	runes1 := []rune(word1)
	runes2 := []rune(word2)
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if runes1[i-1] == runes2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[len1][len2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(minDistance("horse", "ros")) // 3

	println(minDistance("intention", "execution")) // 5

}
