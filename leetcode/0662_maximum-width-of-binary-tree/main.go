package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftPos, rightPos := make([]int, 3000), make([]int, 3000)

	leftPos[0] = 1
	rightPos[0] = 1
	updateLeftRight(root, 1, 1, leftPos, rightPos)

	// lenLeft == lenRight
	lenLeft := len(leftPos)
	result := 0
	for i := 0; i < lenLeft; i++ {
		result = max(result, rightPos[i]-leftPos[i]+1)
	}

	return result
}

func updateLeftRight(node *TreeNode, curPos, depth int, leftPos, rightPos []int) {

	// update leftPos rightPos
	nextLeftPos, nextRightPos := 0, 0
	if node.Left != nil {
		nextLeftPos = curPos*2 - 1
		if leftPos[depth] == 0 || nextLeftPos < leftPos[depth] {
			leftPos[depth] = nextLeftPos
		}
		if rightPos[depth] == 0 || nextLeftPos > rightPos[depth] {
			rightPos[depth] = nextLeftPos
		}
	}
	if node.Right != nil {
		nextRightPos = curPos * 2
		if rightPos[depth] == 0 || nextRightPos > rightPos[depth] {
			rightPos[depth] = nextRightPos
		}
		if leftPos[depth] == 0 || nextRightPos < leftPos[depth] {
			leftPos[depth] = nextRightPos
		}
	}

	// recursion left right node
	if node.Left != nil {
		updateLeftRight(node.Left, nextLeftPos, depth+1, leftPos, rightPos)
	}
	if node.Right != nil {
		updateLeftRight(node.Right, nextRightPos, depth+1, leftPos, rightPos)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(widthOfBinaryTree(&TreeNode{
		2, &TreeNode{
			1, &TreeNode{4, nil, nil}, nil,
		}, &TreeNode{
			4, &TreeNode{5, nil, nil}, nil,
		},
	}))
}
