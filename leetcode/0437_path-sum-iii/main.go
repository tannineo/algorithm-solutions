package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {

	ans := 0

	if root == nil {
		return ans
	}

	prefixSumMap := map[int]int{}

	hasPathSumRecursive(root, 0, sum, &ans, prefixSumMap)

	return ans
}

func hasPathSumRecursive(root *TreeNode, lastSum, target int, ans *int, prefixSumMap map[int]int) {

	if root == nil {
		// do nothing
		return
	}

	// check if there is a path recorded in the map: mappedValue
	// => lastSum + root.Val == mappedValue + target
	if v, ok := prefixSumMap[lastSum+root.Val-target]; ok {
		*ans += v
	}

	if lastSum+root.Val == target { // all the node on path
		*ans++
	}

	// 'map' root.Val
	lastSum += root.Val
	prefixSumMap[lastSum]++

	hasPathSumRecursive(root.Left, lastSum, target, ans, prefixSumMap)
	hasPathSumRecursive(root.Right, lastSum, target, ans, prefixSumMap)

	// 'delete' root.Val from map
	prefixSumMap[lastSum]--
	if prefixSumMap[lastSum] == 0 {
		delete(prefixSumMap, lastSum)
	}
}
