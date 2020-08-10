# 0463_island-perimeter

以水平和垂直两个方向遍历二维数组, 每一行/列的`0, 1`变换就是一条边.

```go
func islandPerimeter(grid [][]int) int {

  if len(grid) == 0 {
    return 0
  }

  if len(grid[0]) == 0 {
    return 0
  }

  result := 0

  // horizontal
  for i := range grid {
    temp := 0
    for j := range grid[i] {
      if grid[i][j] != temp {
        temp = grid[i][j]
        result += 1
      }

      if j == len(grid[i]) - 1 && grid[i][j] == 1 {
        result += 1
      }
    }
  }

  // vertical
  for j := range grid[0] {
    temp := 0
    for i := range grid {
      if grid[i][j] != temp {
        temp = grid[i][j]
        result += 1
      }

      if i == len(grid) - 1 && grid[i][j] == 1 {
        result += 1
      }
    }
  }

  return result
}
```
