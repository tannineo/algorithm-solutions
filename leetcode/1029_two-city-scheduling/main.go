package main

import "sort"

type ByB_ACost [][]int

func (a ByB_ACost) Len() int           { return len(a) }
func (a ByB_ACost) Less(i, j int) bool { return a[i][1]-a[i][0] < a[j][1]-a[j][0] }
func (a ByB_ACost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func twoCitySchedCost(costs [][]int) int {
	sort.Sort(ByB_ACost(costs))

	result := 0
	for i := 0; i < len(costs)/2; i++ {
		result += costs[i][1]
	}
	for i := len(costs) / 2; i < len(costs); i++ {
		result += costs[i][0]
	}

	return result
}

func main() {}
