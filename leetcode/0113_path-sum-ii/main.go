package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {

	ans := [][]int{}

	if root == nil {
		return ans
	}

	pathArr := []int{}

	hasPathSumRecursive(root, 0, sum, &ans, &pathArr)

	return ans
}

func hasPathSumRecursive(root *TreeNode, lastSum, target int, ans *[][]int, pathArr *[]int) {

	if root == nil {
		// when come here, can be sure that: lastSum != target
		return
	}

	if root.Left == nil && root.Right == nil && lastSum+root.Val == target {
		// tricky to copy slice
		cpy := make([]int, len(*pathArr)+1)
		copy(cpy, append(*pathArr, root.Val))
		*ans = append(*ans, cpy)
		return
	}

	// push
	*pathArr = append(*pathArr, root.Val)

	hasPathSumRecursive(root.Left, lastSum+root.Val, target, ans, pathArr)
	hasPathSumRecursive(root.Right, lastSum+root.Val, target, ans, pathArr)

	// pop
	*pathArr = (*pathArr)[0 : len(*pathArr)-1]
}
