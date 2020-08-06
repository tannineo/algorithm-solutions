# 0096_unique-binary-search-trees

第一反应应该是一个数学题.

令`numTrees(0) = 1`

`numTrees(1) = 1`

`numTrees(2) = 2`

`numTrees(3) = 5`

`numTrees(4) = numTrees(0)*numTrees(3) + numTrees(1)*numTrees(2) + numTrees(2)*numTrees(1) + numTrees(3)*numTrees(0) = 5 + 2 + 2 + 5 = 14`

`n`在`[1, 19]`的范围内.

如果输入多, 可以打表.

```go
func numTrees(n int) int {
  arr := make([]int, n+1)
  arr[0] = 1
  arr[1] = 1
  for i := 2; i < n+1; i++ {
    for j := 0; j < i; j++ {
      arr[i] += arr[j]*arr[i-j-1]
    }
  }
  return arr[n]
}
```
