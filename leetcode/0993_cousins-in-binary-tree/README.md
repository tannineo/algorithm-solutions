# 0993_cousins-in-binary-tree

最普适的办法是直接写一个helper去找这个node的depth和parent, 对x, y分别调用.

如果需要考虑剪枝: 不管是dfs还是bfs, 最先找到x或者y, 那么当前深度以下就不要找了.

```go
// isCousins dfs version
func isCousins(root *TreeNode, x int, y int) bool {
  if x == y {
    return false
  }

  limitDepth := -1 // no limit
  depthX := -1 // not found
  depthY := -1 // not found
  parentX := -1 // not found
  parentY := -1 // not found

  dfs(root, x, y, 0, &limitDepth, &depthX, &depthY, &parentX, &parentY)

  // if one of x or y is not found ?
  return depthX != -1 && depthY != -1 && depthX == depthY && parentX != parentY
}

func dfs(root *TreeNode, x, y int, curDepth int, limitDepth *int, depthX, depthY *int, parentX, parentY *int) {
  if *limitDepth != -1 && curDepth > *limitDepth {
    return
  }

  if *depthX != -1 && *depthY != -1 { // found x and y
    return
  }

  if root.Left != nil {
    // find x and y
    if root.Left.Val == x {
      *depthX = curDepth+1
      *parentX = root.Val
      if *limitDepth == -1 {
        *limitDepth = *depthX
      }
    }
    if root.Left.Val == y {
      *depthY = curDepth+1
      *parentY = root.Val
      if *limitDepth == -1 {
        *limitDepth = *depthY
      }
    }
    dfs(root.Left, x, y, curDepth+1, limitDepth, depthX, depthY, parentX, parentY)
  }
  if root.Right != nil {
    // find x and y
    if root.Right.Val == x {
      *depthX = curDepth+1
      *parentX = root.Val
      if *limitDepth == -1 {
        *limitDepth = *depthX
      }
    }
    if root.Right.Val == y {
      *depthY = curDepth+1
      *parentY = root.Val
      if *limitDepth == -1 {
        *limitDepth = *depthY
      }
    }
    dfs(root.Right, x, y, curDepth+1, limitDepth, depthX, depthY, parentX, parentY)
  }
}
```
