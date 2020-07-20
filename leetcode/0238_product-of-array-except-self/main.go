package main

import "fmt"

func productExceptSelf(nums []int) []int {
	lenN := len(nums)

	result := make([]int, lenN)

	// no "zero" cases, since n >= 2

	// 1. product of [0, i-1]
	prod := 1
	result[0] = 1
	for i := 1; i < lenN; i++ {
		prod *= nums[i-1]
		result[i] = prod
	}

	// 2. product of [i+1, n]
	// at the same time, get the result
	prod = 1
	// result[lenN - 1] remains untouched
	for i := lenN - 2; i >= 0; i-- {
		prod *= nums[i+1]
		result[i] *= prod
	}

	return result
}

func main() {
	fmt.Printf("%v\n", productExceptSelf([]int{1, 2, 3, 4})) // [24 12 8 6]
}
