package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	ans := []int{}

	if root == nil {
		return ans
	}

	if root.Left != nil {
		ans = append(ans, inorderTraversal(root.Left)...)
	}

	ans = append(ans, root.Val)

	if root.Right != nil {
		ans = append(ans, inorderTraversal(root.Right)...)
	}

	return ans
}
