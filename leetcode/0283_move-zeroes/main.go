package main

import "fmt"

func moveZeroes(nums []int) {
	lenNums := len(nums)

	if lenNums == 0 {
		return
	}

	right := lenNums - 1

	for right >= 0 && nums[right] == 0 {
		right--
	}
	if right < 0 {
		return // no 0, return
	}

	left := right

	for ; left >= 0; left-- {
		// 1. search for 0
		for left >= 0 && nums[left] != 0 {
			left--
		}
		if left < 0 {
			break // no 0, return
		}

		// println("zero at ", left)

		head := left

		// 2. bubble move 0
		for left+1 < lenNums && nums[left+1] != 0 {
			swapInArray(nums, left, left+1)
			left++
		}

		right++
		left = head
	}
}

func swapInArray(nums []int, a, b int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Printf("%v\n", nums)
}
