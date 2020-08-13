# 0006_zigzag-conversion

模拟去构建.

```go
func convert(s string, numRows int) string {
    if s == "" || numRows == 1 {
        return s
    }

    rows := make([][]rune, numRows)
    for i := range rows {
        rows[i] = []rune{}
    }

    mul := 1
    cur := 0

    runes := []rune(s)

    for _, r := range runes {
        rows[cur] = append(rows[cur], r)

        // change mul
        if cur == numRows-1 {
            mul = -1
        }
        if cur == 0 {
            mul = 1
        }

        // move cur
        cur += mul
    }

    result := ""
    for i := range rows {
        result += string(rows[i])
    }

    return result
}
```
