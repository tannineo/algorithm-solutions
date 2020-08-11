# 0203_remove-linked-list-elements

水.

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    result := head
    for result != nil && result.Val == val {
        result = result.Next
    }
    cur := result
    for cur != nil && cur.Next != nil {
        for cur.Next != nil && cur.Next.Val == val {
            cur.Next = cur.Next.Next
        }
        cur = cur.Next
    }

    return result
}
```
