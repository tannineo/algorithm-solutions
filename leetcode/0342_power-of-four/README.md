# 0342_power-of-four

打表, 或数学方法:

> x is a power of 2 and x % 3 == 1, then x is a power of 4

```text
(2^2k mod 3) = (4^k mod 3) = ((3 + 1)^k mod 3) = 1
((2^2k+1) mod 3) = ((2 * 4^k) mod 3) = ((2 * (3 + 1)^k) mod 3) = 2
```

```go
func isPowerOfFour(num int) bool {
  return num > 0 && num % 3 == 1 && num & (num - 1) == 0
}
```
