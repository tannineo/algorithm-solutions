package main

import "fmt"

func sortColors(nums []int) {
	lenNums := len(nums)
	if lenNums <= 1 {
		return
	}

	// inplace sort
	lastZero := -1
	lastTwo := lenNums
	for i := lastZero + 1; i < lastTwo; i++ {
		for i < lastTwo && nums[i] != 1 {
			if nums[i] == 0 {
				nums[lastZero+1], nums[i] = nums[i], nums[lastZero+1]
				lastZero++
			} else { // == 2
				nums[lastTwo-1], nums[i] = nums[i], nums[lastTwo-1]
				lastTwo--
			}
			if i < lastZero+1 {
				i = lastZero + 1
			}
		}
	}
}

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Printf("%v\n", nums)
}
