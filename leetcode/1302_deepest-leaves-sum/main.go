package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	sum := 0
	sumLevel := -1

	deepestLeavesSumRecursive(root, 0, &sumLevel, &sum)

	return sum
}

func deepestLeavesSumRecursive(root *TreeNode, curLevel int, sumLevel *int, sum *int) {
	if root == nil {
		return
	}

	if curLevel > *sumLevel {
		*sum = root.Val
		*sumLevel = curLevel
	} else if curLevel == *sumLevel {
		*sum += root.Val
	}

	if root.Left != nil {
		deepestLeavesSumRecursive(root.Left, curLevel+1, sumLevel, sum)
	}

	if root.Right != nil {
		deepestLeavesSumRecursive(root.Right, curLevel+1, sumLevel, sum)
	}
}
