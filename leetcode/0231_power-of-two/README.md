# 0231_power-of-two

直觉是将右侧的0去除, 若剩下的结果不是2则为false.

取巧的办法能用两次bitwise操作(外加一次+1), 和一次判断得到答案.

`x & (~x + 1) == x`: (注意`go`的`NOT`是`^n` => `0b111...111 ^ n`, `XOR`是`m^n`)

```go
func isPowerOfTwo(n int) bool {
  if n == 0 {
    return false
  }

  return n & (^n + 1) == n
}
```

`x & (x - 1) == 0`:

```go
func isPowerOfTwo(n int) bool {
  if n == 0 {
    return false
  }

  return n & (n - 1) == 0
}
```
