package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	// This case causes WA
	// if root == nil {
	// 	return sum == 0
	// }

	return hasPathSumRecursive(root, 0, sum)
}

func hasPathSumRecursive(root *TreeNode, lastSum, target int) bool {

	if root == nil {
		// when come here, can be sure that: lastSum != target
		return false
	}

	if root.Left == nil && root.Right == nil && lastSum+root.Val == target {
		return true
	}

	return hasPathSumRecursive(root.Left, lastSum+root.Val, target) || hasPathSumRecursive(root.Right, lastSum+root.Val, target)
}
