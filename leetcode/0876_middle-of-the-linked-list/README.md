# 0876_middle-of-the-linked-list

寻找单向链表中间数, 用两个速率不同的agent去walk这个链表.

纸上模拟一下, 方便确认条件.

```go
func middleNode(head *ListNode) *ListNode {
    fastWalker := head
    slowWalker := head

    count := 0

    for fastWalker.Next != nil {
        fastWalker = fastWalker.Next
        count++

        if count%2 == 1 {
            slowWalker = slowWalker.Next
        }
    }

    return slowWalker
}
```
