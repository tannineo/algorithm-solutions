package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {

	// 0. "zero" cases
	if len(nums) < 2 {
		return 0
	}

	// 1. replace 0 with -1
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}

	// 2. iterate again
	// take record of the sums when they first show up
	sum := 0
	answer := 0
	sumMap := make(map[int]int)
	// 0 as start
	sumMap[0] = -1
	for i, v := range nums {
		sum += v
		if firstPos, ok := sumMap[sum]; ok {
			answer = max(i-firstPos, answer)
		} else {
			sumMap[sum] = i
		}
	}

	return answer
}

func main() {
	println(findMaxLength([]int{0, 1}))                         // 2
	println(findMaxLength([]int{0, 1, 0}))                      // 2
	println(findMaxLength([]int{0, 1, 0, 0, 1}))                // 4
	println(findMaxLength([]int{1, 1, 1, 1, 1, 1, 0, 0, 0, 1})) // 6
}
