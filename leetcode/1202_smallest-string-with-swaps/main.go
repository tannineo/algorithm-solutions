package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Val       int
	Char      byte
	Neighbors []*Node
}

type byBytes []byte

func (bytes byBytes) Len() int           { return len(bytes) }
func (bytes byBytes) Less(i, j int) bool { return bytes[i] < bytes[j] }
func (bytes byBytes) Swap(i, j int)      { bytes[i], bytes[j] = bytes[j], bytes[i] }

func smallestStringWithSwapsSlowVer(s string, pairs [][]int) string {
	// only lowercases a-z
	if len(s) <= 1 {
		return s
	}

	byteS := []byte(s)

	// build graphs by iterating pairs
	// we assume a pair is bi-directional and unique
	mapNodes := map[int]*Node{}
	for _, p := range pairs {
		if _, ok := mapNodes[p[0]]; !ok {
			mapNodes[p[0]] = &Node{Val: p[0], Char: byteS[p[0]], Neighbors: []*Node{}}
		}
		if _, ok := mapNodes[p[1]]; !ok {
			mapNodes[p[1]] = &Node{Val: p[1], Char: byteS[p[1]], Neighbors: []*Node{}}
		}
		mapNodes[p[0]].Neighbors = append(mapNodes[p[0]].Neighbors, mapNodes[p[1]])
		mapNodes[p[1]].Neighbors = append(mapNodes[p[1]].Neighbors, mapNodes[p[0]])
	}

	// dfs to record a group of nodes to swap
	for len(mapNodes) > 0 {
		var startNode *Node
		for _, v := range mapNodes { // get the first node
			startNode = v
			break
		}

		// find connected nodes to swap
		charsToSwap, posToSwap := dfsWithListAndMap(startNode, mapNodes)

		// get the swap result
		sort.Sort(byBytes(charsToSwap))
		sort.Ints(posToSwap)

		// replace
		for i, n := range posToSwap {
			byteS[n] = charsToSwap[i]
		}
	}

	return string(byteS)
}

func dfsWithListAndMap(start *Node, mapNodes map[int]*Node) ([]byte, []int) {
	if start == nil {
		return []byte{}, []int{}
	}
	if _, ok := mapNodes[start.Val]; !ok {
		return []byte{}, []int{}
	}

	charsToSwap := []byte{mapNodes[start.Val].Char}
	posToSwap := []int{start.Val}
	node := mapNodes[start.Val]
	delete(mapNodes, start.Val)

	for _, v := range node.Neighbors {
		if _, ok := mapNodes[v.Val]; !ok {
			continue
		}
		cts, pts := dfsWithListAndMap(v, mapNodes)
		charsToSwap = append(charsToSwap, cts...)
		posToSwap = append(posToSwap, pts...)
	}

	return charsToSwap, posToSwap
}

// ----------------------------------------------------------------

// faster but still TLE
func smallestStringWithSwaps(s string, pairs [][]int) string {
	// only lowercases a-z
	if len(s) <= 1 {
		return s
	}

	// byteS := []byte(s)

	// iterate pairs
	// mapVisited: index -> group of indices
	mapVisited := map[int]*[]int{}
	mapGroupsSet := map[*[]int]bool{}
	for _, p := range pairs {

		// fmt.Println("input: ", p)

		if p[0] == p[1] {
			_, visited := mapVisited[p[0]]
			if !visited {
				newGroup := []int{p[0]}
				mapVisited[p[0]] = &newGroup
				mapGroupsSet[&newGroup] = true
			}

		} else {
			_, visited0 := mapVisited[p[0]]
			_, visited1 := mapVisited[p[1]]

			if !visited0 && !visited1 {
				newGroup := []int{p[0], p[1]}
				mapVisited[p[0]] = &newGroup
				mapVisited[p[1]] = &newGroup
				mapGroupsSet[&newGroup] = true
			} else if visited1 && !visited0 {
				*(mapVisited[p[1]]) = append(*(mapVisited[p[1]]), p[0])
				mapVisited[p[0]] = mapVisited[p[1]]
			} else if visited0 && !visited1 {
				*(mapVisited[p[0]]) = append(*(mapVisited[p[0]]), p[1])
				mapVisited[p[1]] = mapVisited[p[0]]
			} else { // visited0 && visited1
				if mapVisited[p[0]] != mapVisited[p[1]] {
					// combine
					originLen := len(*(mapVisited[p[0]]))
					*(mapVisited[p[0]]) = append(*(mapVisited[p[0]]), (*(mapVisited[p[1]]))...)
					// fmt.Println("newGroup:", newGroup)
					delete(mapGroupsSet, mapVisited[p[1]])
					for i := originLen; i < len(*(mapVisited[p[0]])); i++ {
						mapVisited[(*(mapVisited[p[0]]))[i]] = mapVisited[p[0]]
					}
				}
			}
		}

		// for g := range mapGroupsSet {
		// 	fmt.Println("Group:", *g)
		// }

	}

	// interate recorded group to swap
	for group := range mapGroupsSet {
		charsToSwap := []byte{}
		for _, id := range *group {
			charsToSwap = append(charsToSwap, s[id])
		}

		// get the swap result
		sort.Sort(byBytes(charsToSwap))
		sort.Ints(*group)

		// fmt.Println(*group, "->", string(charsToSwap))

		// replace
		for i, n := range *group {
			s = s[:n] + string(charsToSwap[i]) + s[n+1:]
		}
	}

	return s
}

func main() {
	str := "udyyek"

	fmt.Println(str[len(str):])

	fmt.Println(smallestStringWithSwaps(str, [][]int{
		[]int{3, 3},
		[]int{3, 0},
		[]int{5, 1},
		[]int{3, 1},
		[]int{3, 4},
		[]int{3, 5},
	}))
}
