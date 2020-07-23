package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidSequence(root *TreeNode, arr []int) bool {

	// count nums in arr
	numCounts := make([]int, 10)
	for _, v := range arr {
		numCounts[v]++
	}

	return dfs(root, 1, numCounts, len(arr))
}

func dfs(node *TreeNode, curDepth int, numCounts []int, limit int) bool {

	if node == nil {
		return false
	}

	// check depth
	if curDepth > limit {
		return false
	}

	// check numCounts
	if numCounts[node.Val] == 0 {
		return false
	}
	numCounts[node.Val]--

	if node.Left == nil && node.Right == nil {
		// is leaf
		if curDepth == limit {
			// valid string, check numCounts
			nonZeroFlag := false
			for _, v := range numCounts {
				if v != 0 {
					nonZeroFlag = true
					break
				}
			}
			if !nonZeroFlag {
				// numCounts[node.Val]++ // no need
				return true
			}
		}
	} else {
		// not a leaf
		left := dfs(node.Left, curDepth+1, numCounts, limit)
		if left {
			// numCounts[node.Val]++ // no need
			return true
		}
		right := dfs(node.Right, curDepth+1, numCounts, limit)
		if right {
			// numCounts[node.Val]++ // no need
			return true
		}
	}

	numCounts[node.Val]++
	return false
}

func main() {}
