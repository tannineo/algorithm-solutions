package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderHelper(root *TreeNode, kk *int) int {
	result := 0
	if root.Left != nil {
		result = inOrderHelper(root.Left, kk)
	}

	if result != 0 {
		return result
	}
	*kk--
	if *kk == 0 {
		return root.Val
	}

	if root.Right != nil {
		return inOrderHelper(root.Right, kk)
	}

	return 0
}

func kthSmallest(root *TreeNode, k int) int {
	return inOrderHelper(root, &k)
}

func main() {

}
