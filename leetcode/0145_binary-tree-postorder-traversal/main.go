package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	if root.Left != nil {
		ans = append(ans, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		ans = append(ans, postorderTraversal(root.Right)...)
	}

	ans = append(ans, root.Val)

	return ans
}
