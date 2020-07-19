package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	result, _ := helper(root)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(root *TreeNode) (maxDiameter, maxDepth int) {
	maxDepth = 0
	maxDiameter = 0

	mDiaL, mDepthL := 0, 0
	mDiaR, mDepthR := 0, 0

	// 0. judge "zero" case
	if root == nil || root.Left == nil && root.Right == nil {
		return
	}

	// 1. count the subtree from left and right
	if root.Left != nil {
		mDiaL, mDepthL = helper(root.Left)
		maxDiameter++
	}

	if root.Right != nil {
		mDiaR, mDepthR = helper(root.Right)
		maxDiameter++
	}

	// 2. return the maxDepth and maxDiameter of current part
	maxDepth = max(mDepthL, mDepthR) + 1

	maxDiameter += mDepthL + mDepthR
	maxDiameter = max(max(mDiaL, mDiaR), maxDiameter)

	return
}

func main() {

}
