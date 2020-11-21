package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	// build map from root1
	cachedMap := map[int]struct{}{}

	buildMap(root1, cachedMap)

	return pairWithMap(root2, target, cachedMap)
}

func buildMap(root *TreeNode, cachedMap map[int]struct{}) {
	if root == nil {
		return
	}

	cachedMap[root.Val] = struct{}{}

	buildMap(root.Left, cachedMap)
	buildMap(root.Right, cachedMap)
}

func pairWithMap(root *TreeNode, target int, cachedMap map[int]struct{}) bool {
	if root == nil {
		return false
	}

	leftResult := pairWithMap(root.Left, target, cachedMap)
	if leftResult {
		return true
	}
	rightResult := pairWithMap(root.Right, target, cachedMap)
	if rightResult {
		return true
	}

	_, ok := cachedMap[target-root.Val]
	return ok
}
