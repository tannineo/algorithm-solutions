package main

import "sort"

func eventualSafeNodes(graph [][]int) []int {
	result := []int{}
	if len(graph) == 0 {
		return result
	}

	// build the map: toIndex -> [fromIndex1, fromIndex2, ...]
	// record the nodes without OUT edge
	mapTo2Froms := map[int][]int{}
	mapFrom2Count := map[int]int{}
	results := []int{}
	for fromI, toIndices := range graph {
		if len(toIndices) == 0 {
			results = append(results, fromI)
			continue
		}

		for _, toI := range toIndices {
			if _, ok := mapTo2Froms[toI]; !ok {
				mapTo2Froms[toI] = []int{}
			}
			mapTo2Froms[toI] = append(mapTo2Froms[toI], fromI)
		}
		mapFrom2Count[fromI] = len(toIndices)
	}

	// iterate the results to dfs
	// the discovered "Eventully Safe" points will be added on the right
	for i := 0; i < len(results); i++ {
		for _, fromI := range mapTo2Froms[results[i]] {
			mapFrom2Count[fromI]--
			if mapFrom2Count[fromI] == 0 {
				results = append(results, fromI) // when all the OUT edges are to ES points
			}
		}
	}

	sort.Ints(results)

	return results
}
