# 0171_excel-sheet-column-number

水.

```go
func titleToNumber(s string) int {
    temp := 1
    result := 0
    for i := len(s) - 1; i >= 0; i-- {
        result += (int(s[i] - 'A') + 1) * temp
        temp *= 26
    }

    return result
}
```
