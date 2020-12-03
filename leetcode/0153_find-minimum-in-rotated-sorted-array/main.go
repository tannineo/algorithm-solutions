package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMin(nums []int) int {

	// len(nums) >= 1

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return min(nums[0], nums[1])
	}

	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	mid := len(nums) / 2

	if nums[mid] < nums[0] {
		return findMin(nums[:mid+1])
	}
	return findMin(nums[mid+1:])
}
