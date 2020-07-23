# 0476_number-complement

bitwise operation.

```go
func findComplement(num int) int {
  leftBitOne := 1
  result := 0
  for leftBitOne <= num {
    if leftBitOne & num == 0 {
      result += leftBitOne
    }
    leftBitOne <<= 1
  }

  return result
}
```
