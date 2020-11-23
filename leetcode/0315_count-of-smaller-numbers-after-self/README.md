# 0315_count-of-smaller-numbers-after-self

暴力方法`O(n^2)`, 两重遍历. 约500ms, 能AC.

`O(n*logn)`利用BST(8ms, AC), 从右向左遍历, 构建BST的同时, 每个节点维护左(右)子树及自身相同元素的数量. 这样在构建traversal的时候可以计数.
