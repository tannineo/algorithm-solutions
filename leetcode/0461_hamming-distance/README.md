# 0461_hamming-distance

先`XOR`标记不同的位数, 在遍历获得`1`的个数.

```go
func hammingDistance(x int, y int) int {
    bits := x ^ y

    result := 0

    for bits > 0 {
        result += bits & 1
        bits >>= 1
    }

    return result
}
```

不通过移位, 还有一种bit counting algorithm:

`xor = xor & (xor - 1)`: 移除一个数最右的`1`

http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
