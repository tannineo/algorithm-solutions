package main

// Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if p.Val > q.Val {
		p, q = q, p
	}

	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	var leftNode *TreeNode = nil
	var rightNode *TreeNode = nil
	if root.Val >= p.Val {
		leftNode = lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val <= q.Val {
		rightNode = lowestCommonAncestor(root.Right, p, q)
	}

	if leftNode != nil && rightNode != nil {
		return root
	} else if leftNode != nil {
		return leftNode
	} else if rightNode != nil {
		return rightNode
	}

	return nil
}
