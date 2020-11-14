package main

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	pathArr := [][]int{}
	output := []string{}

	if root == nil {
		return []string{}
	}

	pathNodes := []int{}

	hasPathSumRecursive(root, &pathArr, &pathNodes)

	// convert ans to strings

	for _, v := range pathArr {
		output = append(output, intArr2Str(v))
	}

	return output
}

func hasPathSumRecursive(root *TreeNode, ans *[][]int, pathArr *[]int) {

	if root == nil {
		// when come here, can be sure that: lastSum != target
		return
	}

	if root.Left == nil && root.Right == nil {
		// tricky to copy slice
		cpy := make([]int, len(*pathArr)+1)
		copy(cpy, append(*pathArr, root.Val))
		*ans = append(*ans, cpy)
		return
	}

	// push
	*pathArr = append(*pathArr, root.Val)

	hasPathSumRecursive(root.Left, ans, pathArr)
	hasPathSumRecursive(root.Right, ans, pathArr)

	// pop
	*pathArr = (*pathArr)[0 : len(*pathArr)-1]
}

func intArr2Str(paths []int) string {
	str := ""
	if len(paths) <= 0 {
		return str
	}

	str = strconv.Itoa(paths[0])
	for i := 1; i < len(paths); i++ {
		str += "->" + strconv.Itoa(paths[i])
	}

	return str
}
