package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	}

	newNode := new(TreeNode)

	var t1L, t1R, t2L, t2R *TreeNode
	if t1 != nil {
		t1L, t1R = t1.Left, t1.Right
	}
	if t2 != nil {
		t2L, t2R = t2.Left, t2.Right
	}

	newNode.Left = mergeTrees(t1L, t2L)
	newNode.Right = mergeTrees(t1R, t2R)

	if t1 != nil {
		newNode.Val += t1.Val
	}
	if t2 != nil {
		newNode.Val += t2.Val
	}

	return newNode
}
