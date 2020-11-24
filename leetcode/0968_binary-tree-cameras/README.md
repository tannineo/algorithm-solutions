# 0968_binary-tree-cameras

一个节点有三种可能.

1. 没有被cover.
   - 我们只关心他会被parent的camera给cover的情况, 所以认为他的子节点全被cover(但没有camera).
2. 被子节点cover.
   - 至少有一个子节点有camera, 其他的子节点被cover.
3. 有camera.
   - 子节点状态可能是被cover, 可能是有camera.

所以我们对于一个子树, 返回它上述三种情况的最小camera数, 用于父节点判断.

幸好是二叉树, 判断分支不多...orz

Solution提供了基于dfs的greedy做法.
