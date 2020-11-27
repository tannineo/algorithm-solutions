package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinimumDifference(root *TreeNode) int {
	arr := treeNodesToArr(root)

	ans := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		ans = min(ans, arr[i]-arr[i-1])
	}

	return ans
}

func treeNodesToArr(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	newArr := []int{root.Val}

	left := treeNodesToArr(root.Left)
	right := treeNodesToArr(root.Right)

	newArr = append(left, newArr...)
	newArr = append(newArr, right...)

	return newArr
}
