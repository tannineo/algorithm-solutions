package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestDivisibleSubset(nums []int) []int {

	lenN := len(nums)
	if lenN == 0 {
		return []int{}
	}

	sort.Ints(nums)

	dp := make([]int, lenN)

	for i := 0; i < lenN; i++ {
		dp[i]++
		for j := i + 1; j < lenN; j++ {
			if nums[j]%nums[i] == 0 {
				dp[j] = max(dp[j], dp[i])
			}
		}
	}

	// find max
	maxVal := 0
	at := -1
	for i := 0; i < lenN; i++ {
		if dp[i] > maxVal {
			maxVal = dp[i]
			at = i
		}
	}

	result := []int{nums[at]}

	for i := at; i >= 0; i-- {
		if nums[at]%nums[i] == 0 && dp[at] == (dp[i]+1) {
			result = append(result, nums[i])
			at = i
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", largestDivisibleSubset([]int{1, 2, 3})) // 1, 2 or 1, 3

	fmt.Printf("%v\n", largestDivisibleSubset([]int{1, 2, 4, 8})) // 1, 2, 4, 8

	fmt.Printf("%v\n", largestDivisibleSubset([]int{1, 2, 4, 8, 9, 3, 5, 6})) // 1, 2, 4, 8

	fmt.Printf("%v\n", largestDivisibleSubset([]int{3, 4, 6, 8, 12, 16, 32})) // 4, 8, 16, 32

}
