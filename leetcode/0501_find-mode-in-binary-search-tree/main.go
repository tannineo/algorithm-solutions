package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	answer := []int{}
	curNum, curCount, maxCount := 0, 0, 0

	findModeRecursive(root, &answer, &curNum, &curCount, &maxCount)

	return answer
}

func findModeRecursive(root *TreeNode, answer *[]int, curNum, curCount, maxCount *int) {
	if root == nil {
		return
	}

	// in-order traversal
	findModeRecursive(root.Left, answer, curNum, curCount, maxCount)

	if (*curNum) == root.Val {
		(*curCount)++
	} else { // (*curNum) != root.Val
		(*curNum) = root.Val
		(*curCount) = 1
	}

	if (*curCount) > (*maxCount) {
		(*maxCount) = (*curCount)   // update maxCount
		(*answer) = []int{root.Val} // clear and init a new 'answer'
	} else if (*curCount) == (*maxCount) {
		// add number into answer
		(*answer) = append(*answer, root.Val)
	}

	findModeRecursive(root.Right, answer, curNum, curCount, maxCount)
}
