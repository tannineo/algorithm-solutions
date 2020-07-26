# 0230_kth-smallest-element-in-a-bst

在一个**二叉搜索树(BST)**中寻找从小到大第`k`个元素.

中序遍历, `k-1`发生在中间

## Follow up

> What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

每个节点应该维护以这个节点为根的子树的元素个数, 这样寻找kth smallest和遍历二叉搜索树的过程类似.

insert/delete操作的时候, update这些值.
