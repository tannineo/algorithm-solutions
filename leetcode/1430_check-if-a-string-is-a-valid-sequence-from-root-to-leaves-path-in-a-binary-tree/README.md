# 1430_check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree

题目描述很绕.

我的理解是:

找一个从root到leaf的路径, 这个路径形成的数组是输入数组的重新排列.

注意数字只有1-9, 我们用一个arr维护他们的个数.

DFS, 剪枝策略:

- 记录下行深度, 每次下行时对应数字个数-1:
  - 如果未达到深度就到了叶节点, 退回
  - 如果数字个数扣完, 返回
  - 直到达到arr等长的深度:
    - 如果是叶节点, 检查所有数字是否都为0: 如果是, 整个func返回true
    - 如果不是叶节点, 直接退回
