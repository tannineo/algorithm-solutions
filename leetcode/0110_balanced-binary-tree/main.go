package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	ans, _ := isBalancedWithDepth(root)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isBalancedWithDepth(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftBalanced, leftDepth := isBalancedWithDepth(root.Left)

	if !leftBalanced {
		return false, 0
	}

	rightBalanced, rightDepth := isBalancedWithDepth(root.Right)

	if !rightBalanced {
		return false, 0
	}

	if leftDepth-rightDepth > 1 || rightDepth-leftDepth > 1 {
		return false, 0
	}

	return true, max(leftDepth, rightDepth) + 1
}
