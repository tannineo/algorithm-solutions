package main

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dpCount := make([]int, len(nums))
	dp[0] = 1
	dpCount[0] = 1

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		dpCount[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					dpCount[i] = dpCount[j]
				} else if dp[i] == dp[j]+1 {
					dpCount[i] += dpCount[j]
				}
			}
		}
	}

	// count max
	maxDP := 0
	count := 0

	for i, v := range dp {
		if maxDP < v {
			maxDP = v
			count = dpCount[i]
		} else if maxDP == v {
			count += dpCount[i]
		}
	}

	return count
}

func main() {
	findNumberOfLIS([]int{1, 3, 5, 4, 7})
}
