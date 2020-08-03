# 0700_search-in-a-binary-search-tree

水, 二叉搜索.

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val > val {
        return searchBST(root.Left, val)
    } else if root.Val < val {
        return searchBST(root.Right, val)
    }

    return root
}
```
