package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSum(root *TreeNode) (pathSum int) {
	pathSum, _ = maxPathSumWithMaxSumRoot2Leaf(root)

	return
}

func maxPathSumWithMaxSumRoot2Leaf(root *TreeNode) (pathSum, rootLeafSum int) {
	if root == nil {
		return 0, 0
	}

	if root.Left == nil && root.Right == nil {
		pathSum = root.Val
		if root.Val > 0 {
			rootLeafSum = root.Val
		}
		return
	}

	// calculate only one side
	if root.Left == nil {
		rightPathSum, rightRootLeafSum := maxPathSumWithMaxSumRoot2Leaf(root.Right)
		rootLeafSum = max(root.Val, rightRootLeafSum+root.Val)

		// when calculate pathSum, the left|right RootLeafSum might be cast aside
		if rightRootLeafSum < 0 {
			rightRootLeafSum = 0
		}
		pathSum = max(rightPathSum, rightRootLeafSum+root.Val)
		return
	} else if root.Right == nil {
		leftPathSum, leftRootLeafSum := maxPathSumWithMaxSumRoot2Leaf(root.Left)
		rootLeafSum = max(root.Val, leftRootLeafSum+root.Val)

		// when calculate pathSum, the left|right RootLeafSum might be cast aside
		if leftRootLeafSum < 0 {
			leftRootLeafSum = 0
		}
		pathSum = max(leftPathSum, leftRootLeafSum+root.Val)
		return
	}

	// calculate both side
	leftPathSum, leftRootLeafSum := maxPathSumWithMaxSumRoot2Leaf(root.Left)
	rightPathSum, rightRootLeafSum := maxPathSumWithMaxSumRoot2Leaf(root.Right)

	rootLeafSum = max(root.Val, max(leftRootLeafSum, rightRootLeafSum)+root.Val)

	// when calculate pathSum, the left|right RootLeafSum might be cast aside
	if leftRootLeafSum < 0 {
		leftRootLeafSum = 0
	}
	if rightRootLeafSum < 0 {
		rightRootLeafSum = 0
	}
	pathSum = max(max(leftPathSum, rightPathSum), leftRootLeafSum+rightRootLeafSum+root.Val)
	return
}

func main() {

}
