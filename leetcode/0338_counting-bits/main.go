package main

import "fmt"

func countBits(num int) []int {
	result := []int{0}
	lenR := num + 1

	if lenR == 1 {
		return result
	}

	sqr := 1
	for i := 1; i < lenR; i++ {
		if i >= sqr*2 {
			sqr *= 2
		}
		result = append(result, result[i-sqr]+1)
	}

	return result
}

func main() {
	fmt.Printf("%v\n", countBits(1)) // 0, 1
	fmt.Printf("%v\n", countBits(5)) // 0, 1, 1, 2, 1, 2
	fmt.Printf("%v\n", countBits(8)) // 0, 1, 1, 2, 1, 2, 2, 3, 1
}
