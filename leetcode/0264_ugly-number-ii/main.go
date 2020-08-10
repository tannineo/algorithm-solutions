package main

var dp = []int{1}

func nthUglyNumber(n int) int {
	if len(dp) == 1 {
		initDp()
	}

	return dp[n-1]
}

func initDp() {
	cur2, cur3, cur5 := 0, 0, 0
	for i := 1; i < 1690; i++ {
		temp := min(dp[cur2]*2, dp[cur3]*3, dp[cur5]*5)
		if temp == dp[cur2]*2 {
			cur2++
		}
		if temp == dp[cur3]*3 {
			cur3++
		}
		if temp == dp[cur5]*5 {
			cur5++
		}
		dp = append(dp, temp)
	}
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func main() {
	println(nthUglyNumber(10)) // 12
}
