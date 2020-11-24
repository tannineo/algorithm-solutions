# 0297_serialize-and-deserialize-binary-tree

先/中/后序遍历皆可, deserialize和serialize的过程都是递归traversal, 只是要注意递归返回的`nil`也要被serialize, 作为deserialize的递归结束标志.
