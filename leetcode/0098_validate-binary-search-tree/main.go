package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// testcase [] to be true WTF
	//     if root == nil {
	//         return false
	//     }

	INT64_MAX := int64(^uint64(0) >> 1)
	INT64_MIN := -INT64_MAX - 1

	return isValidBSTRecur(root, INT64_MIN, INT64_MAX)
}

func isValidBSTRecur(root *TreeNode, lb, rb int64) bool {
	if root == nil {
		return true
	}

	rootValInt64 := int64(root.Val)
	if rootValInt64 <= lb || rootValInt64 >= rb {
		return false
	}

	return isValidBSTRecur(root.Left, lb, rootValInt64) && isValidBSTRecur(root.Right, rootValInt64, rb)
}
