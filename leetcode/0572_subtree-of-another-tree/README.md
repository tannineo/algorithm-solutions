# 0572_subtree-of-another-tree

第一反应`O(m*n)`, 遍历`t`, 期间以每一个节点为子树的根, 判断与`s`是否一致.

Solution的解法比较巧妙, 通过加入`left_null`和`right_null`(叶节点的 null children), 并在节点值前加delimiter, 先序遍历得到唯一的一串代表树的字符串, 且能够分辨子树.
