# 0062_unique-paths

第一反应是dp

dp矩阵m x n. dp值是左边和上边的和.

```go
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
    }

    // init value 1
    for i := range dp {
        dp[i][0] = 1
    }
    for j := range dp[0] {
        dp[0][j] = 1
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }

    return dp[m-1][n-1]
}
```

## math

关系到factorial的实现.

https://leetcode.com/problems/unique-paths/solution/
