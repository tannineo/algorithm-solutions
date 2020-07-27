package main

type Node struct {
	Val  int
	Hate []*Node
	Side int // init 0, when partition => +1 || -1
}

func possibleBipartition(N int, dislikes [][]int) bool {
	arrNode := make([]*Node, N+1)

	for _, v := range dislikes {
		// 0. init v[0] v[1]
		if arrNode[v[0]] == nil {
			arrNode[v[0]] = &Node{v[0], []*Node{}, 0}
		}
		if arrNode[v[1]] == nil {
			arrNode[v[1]] = &Node{v[1], []*Node{}, 0}
		}

		// 1. connect v[0] v[1] in bi direction
		arrNode[v[0]].Hate = append(arrNode[v[0]].Hate, arrNode[v[1]])
		arrNode[v[1]].Hate = append(arrNode[v[1]].Hate, arrNode[v[0]])
	}

	// dfs
	for i := 0; i < N+1; i++ {
		if arrNode[i] != nil && arrNode[i].Side == 0 { // dfs not reached
			// start dfs
			dfsArr := [][2]int{[2]int{i, 1}} // [a, b]: node `a` should be at `b` side
			for dfsCur := 0; dfsCur < len(dfsArr); dfsCur++ {
				if arrNode[dfsArr[dfsCur][0]].Side == 0 { // dfs not reached
					arrNode[dfsArr[dfsCur][0]].Side = dfsArr[dfsCur][1]
					for _, v := range arrNode[dfsArr[dfsCur][0]].Hate {
						dfsArr = append(dfsArr, [2]int{v.Val, -dfsArr[dfsCur][1]})
					}
				} else if arrNode[dfsArr[dfsCur][0]].Side != dfsArr[dfsCur][1] {
					return false
				}
			}
		}
	}

	return true
}

func main() {}
