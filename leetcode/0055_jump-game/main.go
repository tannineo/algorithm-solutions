package main

func canJump(nums []int) bool {
	lenNums := len(nums)
	if lenNums == 0 {
		return false
	}
	if lenNums == 1 {
		return true
	}

	dp := make([]bool, lenNums)
	dp[0] = true

	for i, v := range nums {
		if dp[i] { // if can be reached from start
			for j := 1; j <= v && i+j < lenNums; j++ {
				dp[i+j] = true
			}
		}
	}

	return dp[lenNums-1]
}

func main() {
	println(canJump([]int{2, 3, 1, 1, 4})) // true

	println(canJump([]int{3, 2, 1, 0, 4})) // false
}
