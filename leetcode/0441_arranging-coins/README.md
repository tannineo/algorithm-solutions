# 0441_arranging-coins

`(1 + answer) * answer / 2 = n`

=> `0.5 answer^2 + 0.5 answer - n = 0`

用二分法逼近answer, 或者用求根公式: `(-b +- sqrt(b^2 - 4ac)) / 2a`.

- `a = 0.5`
- `b = 0.5`
- `c = - n`

得到答案化为整数. ~~倒推n, 看看倒退结果是否和原来相等.~~

```go
func arrangeCoins(n int) int {
  rows := -0.5 + math.Sqrt(0.25 + 2 * float64(n))
  intRows := int(rows)

  return intRows
}
```
