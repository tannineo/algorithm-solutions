# 0430_flatten-a-multilevel-doubly-linked-list

水, medium也有水题的嗷.

```go
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
func flatten(root *Node) *Node {

  cursor := root

  for {
    for cursor != nil && cursor.Child == nil {
      cursor = cursor.Next
    }

    if cursor == nil {
      return root
    }

    breakBefore := cursor.Next
    childNode := flatten(cursor.Child)
    cursor.Next = childNode
    childNode.Prev = cursor
    cursor.Child = nil

    for cursor.Next != nil {
      cursor = cursor.Next
    }

    cursor.Next = breakBefore
    if breakBefore != nil {
      breakBefore.Prev = cursor
    }
  }
}
```
