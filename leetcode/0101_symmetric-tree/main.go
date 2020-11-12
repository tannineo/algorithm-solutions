package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirrored(root, root)
}

func isMirrored(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isMirrored(left.Left, right.Right) && isMirrored(right.Left, left.Right)
}
