# 0002_add-two-numbers

水. 链表相加, 处理进位.

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var result *ListNode = nil
    cur1, cur2, curR := l1, l2, result
    carry := 0
    for cur1 != nil || cur2 != nil || carry > 0 {
        if curR != nil {
            curR.Next = &ListNode{0, nil}
            curR = curR.Next
        } else {
            result = &ListNode{0, nil}
            curR = result
        }

        if cur1 != nil {
            curR.Val += cur1.Val
            cur1 = cur1.Next
        }
        if cur2 != nil {
            curR.Val += cur2.Val
            cur2 = cur2.Next
        }

        curR.Val += carry
        carry = curR.Val / 10
        curR.Val %= 10
    }

    return result
}
```
