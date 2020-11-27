package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {

	var prev, a, b *TreeNode

	recoverTreeInOrder(root, &prev, &a, &b)

	a.Val, b.Val = b.Val, a.Val
}

func recoverTreeInOrder(root *TreeNode, prev, a, b **TreeNode) {
	if root == nil {
		return
	}

	recoverTreeInOrder(root.Left, prev, a, b)

	if (*prev) != nil {
		if (*prev).Val > root.Val {
			if (*a) == nil {
				*a = (*prev)
				*b = root
			} else {
				*b = root
			}
		}
	}
	(*prev) = root

	recoverTreeInOrder(root.Right, prev, a, b)
}
