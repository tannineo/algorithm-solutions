package main

import "fmt"

func findDuplicates(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		if nums[abs(num)-1] < 0 {
			result = append(result, abs(num))
		}
		nums[abs(num)-1] *= -1
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Printf("%v\n", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})) // 3, 2
}
