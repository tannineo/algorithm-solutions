package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	seq1 := getLeafValSeq(root1)
	seq2 := getLeafValSeq(root2)

	return compareSeq(seq1, seq2)
}

func getLeafValSeq(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	leftSeq := getLeafValSeq(root.Left)
	rightSeq := getLeafValSeq(root.Right)

	return append(leftSeq, rightSeq...)
}

func compareSeq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
