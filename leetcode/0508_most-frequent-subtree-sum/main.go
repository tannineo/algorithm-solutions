package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	sumMap := map[int]int{}
	recordFTSWithMap(root, sumMap)

	// find answer(s)
	anss := []int{}
	maxSum := 0
	for k, v := range sumMap {
		if v > maxSum {
			anss = []int{k}
			maxSum = v
		} else if v == maxSum {
			anss = append(anss, k)
		}
	}

	return anss
}

func recordFTSWithMap(root *TreeNode, sumMap map[int]int) int {

	// no-nil root ensured

	leftSum, rightSum := 0, 0

	if root.Left != nil {
		leftSum = recordFTSWithMap(root.Left, sumMap)
	}
	if root.Right != nil {
		rightSum = recordFTSWithMap(root.Right, sumMap)
	}

	sum := root.Val + leftSum + rightSum
	sumMap[sum]++
	return sum
}
