package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := &[][]int{}

	levelOrderTraversal(root, 1, ans)

	return *ans
}

func levelOrderTraversal(root *TreeNode, nextLevel int, ans *[][]int) {
	if root == nil {
		return
	}

	if len(*ans) < nextLevel {
		*ans = append(*ans, []int{})
	}

	(*ans)[nextLevel-1] = append((*ans)[nextLevel-1], root.Val)

	if root.Left != nil {
		levelOrderTraversal(root.Left, nextLevel+1, ans)
	}

	if root.Right != nil {
		levelOrderTraversal(root.Right, nextLevel+1, ans)
	}

}
