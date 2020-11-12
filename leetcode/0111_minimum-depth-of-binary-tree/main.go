package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	if leftDepth == 0 && rightDepth == 0 {
		return 1
	} else if leftDepth == 0 {
		return rightDepth + 1
	} else if rightDepth == 0 {
		return leftDepth + 1
	}

	if leftDepth < rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
