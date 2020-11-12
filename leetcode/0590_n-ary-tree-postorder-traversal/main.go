package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	for _, n := range root.Children {
		ans = append(ans, postorder(n)...)
	}

	ans = append(ans, root.Val)

	return ans
}
