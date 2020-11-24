# 0968_binary-tree-cameras

一个节点有三种可能.

1. 没有被cover.
   - 我们只关心他会被parent的camera给cover的情况, 所以认为他的子节点全被cover(但没有camera).
2. 被子节点cover.
   - 至少有一个子节点有camera, 其他的子节点被cover.
3. 有camera.
   - 子节点状态可能是三种情况之一.

注意的是三种情况并不会同时出现在一个节点上, 对于一个非空节点, 情况1和情况2可能不会存在, 判断时请注意.

我们对于一个子树, 返回它上述三种情况的最小camera数, 用于父节点判断.

幸好是二叉树, 判断分支不多...orz

Solution提供了基于dfs的greedy做法.
