package main

import "math"

type City struct {
	Val   int
	Conn  []*City
	Costs []int

	IfRecorded bool
	IfCanReach bool
	RecordStep int
	RecordCost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if src == dst {
		return 0
	}

	result := -1

	// 1. construct graph
	cities := make([]*City, n)
	for _, v := range flights {
		if cities[v[0]] == nil {
			cities[v[0]] = &City{v[0], []*City{}, []int{}, false, false, 0, math.MaxInt32}
		}
		if cities[v[1]] == nil {
			cities[v[1]] = &City{v[1], []*City{}, []int{}, false, false, 0, math.MaxInt32}
		}
		cities[v[0]].Conn = append(cities[v[0]].Conn, cities[v[1]])
		cities[v[0]].Costs = append(cities[v[0]].Costs, v[2])
	}

	// 2. dfs
	dfsHelper(cities[src], cities[dst], 0, &result, K)

	return result
}

func dfsHelper(from *City, to *City, totalCost int, savedResult *int, leftJump int) bool {
	if from.Val == to.Val {
		// check cost
		if *savedResult == -1 {
			*savedResult = totalCost
		} else if *savedResult > totalCost {
			*savedResult = totalCost
		}
		return true
	}

	if leftJump < 0 {
		return false
	}

	if from.IfRecorded {
		if from.IfCanReach && totalCost >= from.RecordCost && from.RecordStep >= leftJump {
			return false
		} else if !from.IfCanReach && from.RecordStep >= leftJump {
			return false
		}
	}

	from.IfRecorded = true

	lenConn := len(from.Conn)
	flag := false
	for i := 0; i < lenConn; i++ {
		result := dfsHelper(from.Conn[i], to, totalCost+from.Costs[i], savedResult, leftJump-1)
		flag = flag || result
		if from.IfCanReach && result {
			// update record
			if totalCost <= from.RecordCost && leftJump >= from.RecordStep {
				from.RecordCost = totalCost
				from.RecordStep = leftJump
			}
		} else if !from.IfCanReach && result {
			from.IfCanReach = true
			from.RecordCost = totalCost
			from.RecordStep = leftJump
		} else if !from.IfCanReach && !result {
			if from.RecordStep > leftJump {
				from.RecordStep = leftJump
			}
		}
	}

	return flag
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {}
