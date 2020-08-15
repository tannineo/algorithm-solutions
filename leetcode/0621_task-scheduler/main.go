package main

import "sort"

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	// stack tasks by type
	taskPile := make([]int, 26)
	total := 0
	for _, t := range tasks {
		taskPile[int(t-'A')]--
		total++
	}

	// sort
	sort.Ints(taskPile)
	maxC := taskPile[0]
	count := 0
	for _, tc := range taskPile {
		if tc == maxC {
			count++
		}
	}

	possible := (maxC-1)*(n+1) + count

	return max(total, possible)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {}
