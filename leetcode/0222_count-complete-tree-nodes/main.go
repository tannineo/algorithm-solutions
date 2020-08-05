package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	CBTcount, _, _, _ := countCBT(root)

	return CBTcount
}

func countCBT(root *TreeNode) (CBTcount int, isCBT bool, depth int, isFullBiTree bool) {
	if root == nil {
		return 0, false, 0, false
	}

	if root.Left == nil && root.Right == nil {
		return 1, true, 1, true
	}

	leftCBTCount, leftIsCBT, leftDepth, leftIsFullBiTree := countCBT(root.Left)
	rightCBTCount, rightIsCBT, rightDepth, rightIsFullBiTree := countCBT(root.Right)

	if leftDepth == rightDepth && leftIsFullBiTree && rightIsCBT {
		CBTcount++
		isCBT = true
	} else if leftDepth == rightDepth+1 && leftIsCBT && rightIsFullBiTree {
		CBTcount++
		isCBT = true
	} else if leftDepth == 1 && rightDepth == 0 {
		CBTcount++
		isCBT = true
	}

	CBTcount += leftCBTCount + rightCBTCount
	depth = max(leftDepth, rightDepth) + 1

	// is full bi-tree
	isFullBiTree = leftIsFullBiTree && rightIsFullBiTree && leftDepth == rightDepth

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {}
