# 0275_h-index-ii

Hmmm... medium难度??

已经排好序了, 从citation多的一头数.

```go
func hIndex(citations []int) int {
    lenC := len(citations)

    if lenC == 0 {
        return 0
    }

    result := 0

    for i := lenC - 1; i >= 0; i-- {
        if citations[i] >= result+1 {
            result++
        }
    }

    return result
}
```
