package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	mid := len(nums) / 2

	left := sortedArrayToBST(nums[:mid])
	right := sortedArrayToBST(nums[mid+1:])

	newNode := TreeNode{nums[mid], left, right}

	return &newNode
}

func main() {
	sortedArrayToBST([]int{0, 1, 2})
}
