package main

// Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)

	if leftNode != nil && rightNode != nil {
		return root
	} else if leftNode != nil {
		return leftNode
	} else if rightNode != nil {
		return rightNode
	}

	return nil
}
