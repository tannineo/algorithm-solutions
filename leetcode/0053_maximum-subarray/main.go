package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max_sum := nums[0]
	last_val := nums[0]

	for i, v := range nums {
		if i == 0 {
			continue
		}

		last_val = max(last_val+v, v)
		max_sum = max(max_sum, last_val)
	}

	return max_sum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
