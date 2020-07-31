package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)

	if right == 0 {
		return 0
	}

	var guess int
	for left < right {
		guess = (left + right) / 2
		if nums[guess] == target {
			return guess
		}

		if nums[guess] < target {
			left = guess + 1
		} else { // nums[guess] > target
			right = guess
		}
	}

	if nums[guess] > target {
		return guess
	}
	// nums[guess]  < target
	return guess + 1
}

func main() {
	println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4
}
