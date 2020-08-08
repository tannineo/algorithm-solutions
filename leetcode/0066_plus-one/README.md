# 0066_plus-one

水. 注意, `digits`代表的数必`>0`.

```go
func plusOne(digits []int) []int {
  cur := len(digits) - 1
  carry := 0
  for {
    digits[cur] += 1
    carry = digits[cur] / 10
    digits[cur] %= 10

    if carry == 0 {
      break
    }

    cur--

    if cur == -1 { // to the left most
      digits = append([]int{1}, digits...)
      break
    }
  }

  return digits
}
```
