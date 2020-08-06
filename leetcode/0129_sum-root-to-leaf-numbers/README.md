# 0129_sum-root-to-leaf-numbers

遍历的时候, 从上到下传递数字.

```go
func sumNumbers(root *TreeNode) int {
  result := 0

  sumRoot2Leaf(root, 0, &result)

  return result
}

func sumRoot2Leaf(node *TreeNode, num int, result *int) {
  if node == nil {
    return
  }

  num = 10 * num + node.Val

  if node.Left == nil && node.Right == nil {
    *result += num
    return
  }

  sumRoot2Leaf(node.Left, num, result)
  sumRoot2Leaf(node.Right, num, result)
}
```
