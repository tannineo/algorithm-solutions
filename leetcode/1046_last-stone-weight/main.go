package main

import (
	"container/heap"
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

type StoneHeap []int

// sort.Interface
func (s StoneHeap) Len() int           { return len(s) }
func (s StoneHeap) Less(i, j int) bool { return s[i] > s[j] }
func (s StoneHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// heap.Interface
func (s *StoneHeap) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *StoneHeap) Pop() interface{} {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func lastStoneWeightWithHeap(stones []int) int {
	sh := StoneHeap{}

	heap.Init(&sh)

	for _, v := range stones {
		heap.Push(&sh, v)
	}

	for len(sh) > 1 {
		stone1 := heap.Pop(&sh)
		stone2 := heap.Pop(&sh)
		rest := stone1.(int) - stone2.(int)
		if rest > 0 {
			heap.Push(&sh, rest)
		}
	}

	if len(sh) == 0 {
		return 0
	}
	return heap.Pop(&sh).(int)
}

func main() {
	println(lastStoneWeightWithHeap([]int{2, 7, 4, 1, 8, 1})) // 1

	println(lastStoneWeightWithHeap([]int{3, 7, 8})) // 2
}
