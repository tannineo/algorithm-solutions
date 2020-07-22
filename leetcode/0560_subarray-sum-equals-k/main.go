package main

func subarraySum(nums []int, k int) int {
	result := 0

	lenNums := len(nums)
	l := 1
	for l <= lenNums {
		// iterate over nums using a slide window [left, right)

		// 0. init slide window
		left := 0
		right := left + l
		sum := 0
		for _, v := range nums[left:right] {
			sum += v
		}
		// judge the inital value
		if sum == k {
			result++
		}

		// 1. slide the window
		for right < lenNums {
			sum -= nums[left]
			left++
			sum += nums[right]
			right++

			if sum == k {
				result++
			}
		}

		l++
	}

	return result
}

func subarraySumMap(nums []int, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1

	result := 0

	sum := 0

	for _, v := range nums {
		sum += v

		if count, ok := sumMap[sum-k]; ok {
			result += count
		}

		sumMap[sum]++

	}

	return result
}

func main() {
	println(subarraySum([]int{1, 1, 1}, 2)) // 2

	println(subarraySumMap([]int{1, 1, 1}, 2)) // 2
}
