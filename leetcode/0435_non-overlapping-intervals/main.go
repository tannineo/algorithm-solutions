package main

import "sort"

type ByEnd [][]int

func (be ByEnd) Less(a, b int) bool {
	return be[a][1] < be[b][1]
}

func (be ByEnd) Swap(a, b int) {
	be[a], be[b] = be[b], be[a]
}

func (be ByEnd) Len() int {
	return len(be)
}

func eraseOverlapIntervals(intervals [][]int) int {

	lenI := len(intervals)

	if lenI <= 1 {
		return 0
	}

	// sort before dp
	sort.Sort(ByEnd(intervals))
	count := 1
	end := intervals[0][1]

	for i := 1; i < lenI; i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}

	return lenI - count
}

func main() {
	println(eraseOverlapIntervals([][]int{
		{1, 100},
		{11, 22},
		{1, 11},
		{2, 12},
	}))
}
