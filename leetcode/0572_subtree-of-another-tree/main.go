package main

import (
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	sStr := tree2string(s)
	tStr := tree2string(t)

	return strings.Contains(sStr, tStr)
}

func tree2string(root *TreeNode) string {
	// non-empty ensured

	str := "#" + strconv.Itoa(root.Val)

	if root.Left == nil {
		str += "#LN"
	} else {
		str += tree2string(root.Left)
	}

	if root.Right == nil {
		str += "#RN"
	} else {
		str += tree2string(root.Right)
	}

	return str
}
