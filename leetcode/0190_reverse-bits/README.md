# 0190_reverse-bits

```go
func reverseBits(num uint32) uint32 {
    result := uint32(0)
    currentBit := uint32(1)

    for i := 0; i < 32; i++ {
        result <<= 1
        if num & currentBit > 0 {
            result |= 1
        }
        currentBit <<= 1
    }

    return result
}
```
