# 1428_leftmost-column-with-at-least-a-one

按照题目描述, 这个`BinaryMatrix`里的`0`都挤在左侧, `1`都挤在右侧.

我们可以从第一行的边界向下找, 每次更新边界的"最左值".

题目要求在`1000`次操作中完成, 这样的操作应该不会超过`300`(100*100矩阵).

```go
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
  dim := binaryMatrix.Dimensions()
  row, column := dim[0], dim[1]

  rightLeast0 := -1

  // init
  for j := 0; j < column && binaryMatrix.Get(0, j) == 0; j++ {
    rightLeast0 = j
  }
  if rightLeast0 == -1 {
    return 0
  }

  // walk to find least zero pos
  for i := 1; i < row; i++ {
    if binaryMatrix.Get(i, rightLeast0) == 1 {
      for rightLeast0 >= 0 && binaryMatrix.Get(i, rightLeast0) == 1 {
        rightLeast0--
      }
      if rightLeast0 == -1 {
        return 0
      }
    }
  }

  if rightLeast0 == column - 1 {
    return -1
  }

  return rightLeast0 + 1
}
```
