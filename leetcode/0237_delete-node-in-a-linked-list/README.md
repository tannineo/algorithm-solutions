# 0237_delete-node-in-a-linked-list

水到爆炸.

题目怪到一开始golang的函数签名我看不懂.

输入是指向某个节点的指针, 仅此而已.

这个节点就是你要删去的点...

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
  node.Val, node.Next = node.Next.Val, node.Next.Next
}
```
