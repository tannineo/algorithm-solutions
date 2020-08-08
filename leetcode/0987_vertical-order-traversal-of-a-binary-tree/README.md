# 0987_vertical-order-traversal-of-a-binary-tree

遍历, 然后按照自定规则排序, 最后重建数组.

```go
type ValNode struct {
  Val int
  X int
  Y int
}

type ByYXVal []ValNode

func (n ByYXVal) Len() int { return len(n) }
func (n ByYXVal) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n ByYXVal) Less(i, j int) bool {
  if n[i].X < n[j].X {
    return true
  }
  if n[i].X == n[j].X && n[i].Y > n[j].Y { // Y: 0 -> -Inf
    return true
  }
  if  n[i].X == n[j].X && n[i].Y == n[j].Y && n[i].Val < n[j].Val { // smaller first
    return true
  }
  return false
}

func verticalTraversal(root *TreeNode) [][]int {
  valNodes := []ValNode{}

  // traversal
  addNode(root, 0, 0, &valNodes)

  // sort
  sort.Sort(ByYXVal(valNodes))

  // build answer arrays
  answer := [][]int{}
  curX := -1 << 31
  for _, node := range valNodes {
    if node.X != curX {
      // append
      answer = append(answer, []int{})
      curX = node.X
    }
    lenA := len(answer)
    answer[lenA-1] = append(answer[lenA-1], node.Val)
  }

  return answer
}

func addNode(node *TreeNode, x, y int, valNodes *[]ValNode) {
  if node == nil {
    return
  }

  *valNodes = append(*valNodes, ValNode{node.Val, x, y})

  addNode(node.Left, x-1, y-1, valNodes)
  addNode(node.Right, x+1, y-1, valNodes)
}
```
