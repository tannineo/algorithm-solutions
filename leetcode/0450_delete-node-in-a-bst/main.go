package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	// root.Val == key
	if root.Left == nil && root.Right == nil {
		return nil
	} else if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	}

	// find predecessor
	predParent := root
	pred := root.Left

	if pred.Right == nil { // just one step to the left?
		pred.Right = root.Right
		return pred
	}

	for pred.Right != nil {
		predParent = pred
		pred = pred.Right
	}

	if pred.Left == nil && pred.Right == nil { // is leaf?
		predParent.Right = nil
		pred.Left = root.Left
		pred.Right = root.Right
		return pred
	}

	root.Val = pred.Val                                       // swap value
	predParent.Right = deleteNode(predParent.Right, pred.Val) // recursive deletion

	return root
}
