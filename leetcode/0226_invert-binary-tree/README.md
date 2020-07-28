# 0226_invert-binary-tree

BST 遍历

> Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.

```go
func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return nil
  }

  root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

  return root
}
```
