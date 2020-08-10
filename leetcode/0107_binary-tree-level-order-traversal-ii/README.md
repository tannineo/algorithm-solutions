# 0107_binary-tree-level-order-traversal-ii

preorder traversal.

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
  result := [][]int{}

  preorderHelper(root, 0, &result)

  reverseSlice(result)

  return result
}

func preorderHelper(node *TreeNode, depth int, result *[][]int) {
  if node == nil {
    return
  }

  for len(*result) <= depth {
    *result = append(*result, []int{})
  }

    (*result)[depth] = append((*result)[depth], node.Val)

  preorderHelper(node.Left, depth+1, result)
  preorderHelper(node.Right, depth+1, result)
}

func reverseSlice(s [][]int) {
  i, j := 0, len(s)-1
  for i < j {
    s[i], s[j] = s[j], s[i]
    i++
    j--
  }
}
```
