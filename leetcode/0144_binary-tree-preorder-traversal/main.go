package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)

	if root.Left != nil {
		ans = append(ans, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		ans = append(ans, preorderTraversal(root.Right)...)
	}

	return ans
}
