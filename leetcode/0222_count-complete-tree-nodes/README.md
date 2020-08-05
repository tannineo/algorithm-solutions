# 0222_count-complete-tree-nodes

递归, 后序遍历.

当前点是否为complete binary tree(CBT), 取决于左右子节点的答案.

- 一个节点无左右子节点为一个CBT.
- 一个节点仅有左节点
  - 且左节点是叶子节点(深度为1), 为CBT
- 一个节点有左右两个子节点
  - 且左右节点深度一致
  - 且左节点是满二叉树(也是CBT), 右节点是CBT,
  - 或
  - 且左节点深度 - 1 = 右节点深度
  - 且左节点是CBT, 右节点是满二叉树(也是CBT),
  - 为CBT
- 其余所有情况都不是CBT

所以遍历的时候我们需要返回三个值:

- 该节点是否为CBT
- 以该节点为root的子树深度
- 以该节点是否为满二叉树

**但是**, 上述是通用方法, 题目给出的树自身已经是CBT!!!!!!!!!

其实就是数节点个数.

```go
func countNodes(root *TreeNode) int {
  if root == nil {
    return 0
  }

  return 1 + countNodes(root.Left) + countNodes(root.Right)
}
```
