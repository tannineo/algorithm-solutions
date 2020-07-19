# 0543_diameter-of-binary-tree

基于二叉树的后序遍历(`post-order traversal`, 相对的有`pre-order`, `in-order`).

对于每个subtree, 我们需要返回两个值:

- subtree的深度`maxDepth`
- 符合题意的答案`maxDiameter`
  - `maxDiameter`是以下三个值的最大值:
    - 左右两边深度的和`+2`(一边没有child就要`-1`)
    - 左边的subtree的`maxDiameter`
    - 左边的subtree的`maxDiameter`
