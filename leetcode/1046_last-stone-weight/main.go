package main

import (
	"sort"
)

func bubbleSort(arr []int, newVal int) []int {
	for i, v := range arr {
		if v > newVal {
			tempArr := make([]int, len(arr[i:]))
			copy(tempArr, arr[i:])
			arr = append(arr[:i], newVal)
			arr = append(arr, tempArr...)
			return arr
		}
	}
	arr = append(arr, newVal)
	return arr
}

func lastStoneWeight(stones []int) (result int) {

	result = 0
	if len(stones) == 0 {
		return
	}

	// 0. sort all first
	sort.Ints(stones) // 0, 1, 2, 3

	for len(stones) > 0 {
		lenS := len(stones)
		// 0.1 if only 1 stone left
		if lenS == 1 {
			return stones[0]
		}

		// 1. find 2 stones to smash
		if stones[lenS-1] == stones[lenS-2] {
			stones = stones[:lenS-2]
			// need no sort here
		} else {
			newStone := stones[lenS-1] - stones[lenS-2]
			stones = stones[:lenS-2]
			stones = bubbleSort(stones, newStone)
		}
	}

	return
}

func main() {
	println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) // 1

	println(lastStoneWeight([]int{3, 7, 8})) // 2
}
