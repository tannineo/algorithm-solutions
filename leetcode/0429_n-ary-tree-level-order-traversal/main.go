package main

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ans := [][]int{}

	levelOrderTraversal(root, 0, &ans)

	return ans
}

func levelOrderTraversal(root *Node, level int, store *[][]int) {

	if root == nil {
		return
	}

	if len(*store) <= level {
		*store = append(*store, []int{})
	}

	if (*store)[level] == nil {
		(*store)[level] = []int{}
	}

	(*store)[level] = append((*store)[level], root.Val)

	for _, v := range root.Children {
		levelOrderTraversal(v, level+1, store)
	}
}
