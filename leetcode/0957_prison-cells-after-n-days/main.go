package main

import "fmt"

func prisonAfterNDays(cells []int, N int) []int {
	if N == 0 { // [1, 10^9]
		return cells
	}

	// 0. initialize number for cells
	num := arr2Int(cells)

	// 1. simulate
	for {
		num = getNext(num)
		N--
		N %= 14
		if N <= 0 {
			break
		}
	}

	return int2Arr(num)
}

func getNext(num int) int {
	return  (^((num << 1) ^ (num >> 1))) & 0b01111110
}

func arr2Int(arr []int) int {
	num := 0
	for _, bit := range arr {
		num <<= 1
		num |= bit
	}
	return num
}

func int2Arr(num int) []int {
	result := []int{}
	cursor := 0b10000000
	for cursor > 0 {
		if num&cursor > 0 {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
		cursor >>= 1
	}
	return result
}

func findCircle() []int {
	result := []int{}
	for i := 0 ; i <= 0b01111110; i++ {
		if i & 0b01111110 != i {
			result = append(result, 0)
			continue
		}
		println(i)
		num := i
		step := 0
		for {
			num = getNext(num)
			step++
			if num == i {
				break;
			}
		}
		result = append(result, step)
	}

	return result
}

func main() {

	// fmt.Printf("%v\n", findCircle())

	fmt.Printf("%v\n", prisonAfterNDays([]int{0, 1, 0, 1, 1, 0, 0, 1}, 7)) // [0, 0, 1, 1, 0, 0, 0, 0]

	fmt.Printf("%v\n", prisonAfterNDays([]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000)) // 0,0,1,1,1,1,1,0

	fmt.Printf("%v\n", prisonAfterNDays([]int{0,1,1,1,0,0,0,0}, 99)) //
}
