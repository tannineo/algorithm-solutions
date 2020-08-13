# 0070_climbing-stairs

dp. 其实是Fibonacci数列.

```go
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }

    first, second := 1, 1

    for i := 2; i <= n; i++ {
        first, second = second, first + second
    }

    return second
}
```

Fibonacci有直接计算的公式, 见[solution](https://leetcode.com/problems/climbing-stairs/solution/)
