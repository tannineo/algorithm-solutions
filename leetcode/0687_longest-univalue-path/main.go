package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, _ := lupRecur(root)
	// we count the nodes
	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lupRecur(root *TreeNode) (answer, pathFromRoot int) {
	if root == nil {
		return 0, 0
	}

	if root.Left == nil && root.Right == nil {
		return 1, 1
	}

	leftAns, leftPFR := lupRecur(root.Left)
	rightAns, rightPFR := lupRecur(root.Right)

	// pathFromRoot

	leftPath, rightPath := 1, 1

	if root.Left != nil && root.Val == root.Left.Val {
		leftPath += leftPFR
	}

	if root.Right != nil && root.Val == root.Right.Val {
		rightPath += rightPFR
	}

	pathFromRoot = max(leftPath, rightPath)

	// answer

	answer = max(leftAns, rightAns)
	answer = max(answer, pathFromRoot)

	if root.Left != nil && root.Right != nil && root.Val == root.Left.Val && root.Val == root.Right.Val {
		answer = max(answer, leftPFR+rightPFR+1)
	}

	return
}
