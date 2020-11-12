package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	ans = append(ans, root.Val)

	for _, n := range root.Children {
		ans = append(ans, preorder(n)...)
	}

	return ans
}
