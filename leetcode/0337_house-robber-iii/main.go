package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	containSelf, justChildren := robOptions(root)

	return max(containSelf, justChildren)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robOptions(root *TreeNode) (containSelf, justChildren int) {
	if root == nil {
		return 0, 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val, 0
	}

	leftCS, leftJC := robOptions(root.Left)
	rightCS, rightJC := robOptions(root.Right)

	containSelf = root.Val + leftJC + rightJC
	justChildren = max(leftCS, leftJC) + max(rightCS, rightJC)
	return
}
