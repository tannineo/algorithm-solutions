package main

import (
	"fmt"
	"sort"
)

type ByHeightThenPeopleBefore [][]int

func (b ByHeightThenPeopleBefore) Len() int { return len(b) }
func (b ByHeightThenPeopleBefore) Less(i, j int) bool {
	if b[i][0] > b[j][0] {
		return true
	} else if b[i][0] < b[j][0] {
		return false
	}
	// ==
	if b[i][1] < b[j][1] {
		return true
	}
	return false
}
func (b ByHeightThenPeopleBefore) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func reconstructQueue(people [][]int) [][]int {

	lenP := len(people)
	if lenP == 0 {
		return people
	}

	sort.Sort(ByHeightThenPeopleBefore(people))

	// inplace rearrange
	for i := 1; i < lenP; i++ {
		if people[i][1] < i {
			curPos := i
			for curPos > people[curPos][1] {
				people[curPos], people[curPos-1] = people[curPos-1], people[curPos] // swap
				curPos--
			}
		} // '== i' do nothing, '> i' won't exist
	}

	return people
}

func main() {
	fmt.Printf("%v\n", reconstructQueue([][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	})) // [5,0], [7,0], [5,2], [6,1], [4,4], [7,1]
}
