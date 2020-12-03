package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	startVal := node.Val

	// dfs
	mapVisited := map[int][]int{}
	dfsWithMap(node, mapVisited)

	// rebuild
	rebuiltNodes := map[int]*Node{}
	for k := range mapVisited {
		rebuiltNodes[k] = &Node{Val: k}
	}
	for k := range rebuiltNodes {
		for _, con := range mapVisited[k] {
			rebuiltNodes[k].Neighbors = append(rebuiltNodes[k].Neighbors, rebuiltNodes[con])
		}
	}

	return rebuiltNodes[startVal]
}

func dfsWithMap(node *Node, mapVisited map[int][]int) {
	if node == nil {
		return
	}

	if _, ok := mapVisited[node.Val]; ok {
		return
	}

	nArr := []int{}
	for _, v := range node.Neighbors {
		nArr = append(nArr, v.Val)
	}
	mapVisited[node.Val] = nArr

	for _, v := range node.Neighbors {
		dfsWithMap(v, mapVisited)
	}
}
