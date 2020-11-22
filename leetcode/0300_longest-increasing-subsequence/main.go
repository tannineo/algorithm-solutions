package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(n^2)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	for i, cur := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < cur {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	// find max
	answer := 1
	for _, v := range dp {
		answer = max(answer, v)
	}

	return answer
}

// O(n*logn)
func lengthOfLISBS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	ans := 1
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		pos := bs(dp, 0, ans-1, nums[i])
		dp[pos] = nums[i]
		if pos == ans {
			ans++
		}
	}
	return ans
}

// bs finds the insert index of `target` in arr[left, right]
func bs(arr []int, left, right, target int) int {
	for left <= right {
		cur := (right-left)/2 + left

		if arr[cur] < target {
			left = cur + 1
		} else if arr[cur] > target {
			right = cur - 1
		} else {
			return cur
		}
	}

	return left
}

func main() {
	// fmt.Println(lengthOfLISBS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4

	fmt.Println(lengthOfLISBS([]int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12})) // 6
}
