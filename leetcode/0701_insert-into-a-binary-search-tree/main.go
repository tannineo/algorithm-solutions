package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil { // new tree
		return &TreeNode{val, nil, nil}
	}

	// unique val guaranteed
	if root.Val > val {
		if root.Left == nil {
			root.Left = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Left, val)
		}
	} else if root.Val < val {
		if root.Right == nil {
			root.Right = &TreeNode{val, nil, nil}
		} else {
			insertIntoBST(root.Right, val)
		}
	}

	return root
}
