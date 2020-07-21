package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	lenP := len(preorder)
	stack := make([]*TreeNode, lenP)
	top := -1
	root := &TreeNode{}

	for i, v := range preorder {
		if i == 0 {
			// init
			root.Val = v
			top = 0
			stack[top] = root
			continue
		}

		curNode := stack[top]

		if stack[top].Val < v {
			for top >= 0 && stack[top].Val < v {
				curNode = stack[top]
				top--
			}
			curNode.Right = &TreeNode{Val: v}
			top++
			stack[top] = curNode.Right
		} else { // v < stack[top].Val
			curNode.Left = &TreeNode{Val: v}
			top++
			stack[top] = curNode.Left
		}

	}

	return root
}

func main() {

	bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
}
