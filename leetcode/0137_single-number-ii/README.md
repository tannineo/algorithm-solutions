# 0137_single-number-ii

位运算的神奇用法...

[0136_single-number](../0136_single-number) 是"其余元素只出现两次"的版本. 一个`XOR`从头到尾就能搞定.

1. 第一次出现, `seenOnce`记录, 且因为`seenOnce`记录了, `seenTwice`不会记录
2. 第二次出现, 从`seenOnce`中删除, 向`seenTwice`中添加记录
3. 第三次出现, 因为`seenTwice`中记录了, `seenOnce`中不会记录, 再从`seenTwice`中删除

```go
func singleNumber(nums []int) int {
  seenOnce, seenTwice := 0, 0

  for _, num := range nums {
    seenOnce = ^seenTwice & (seenOnce ^ num)
    seenTwice = ^seenOnce & (seenTwice ^ num)
  }

  return seenOnce
}
```

## 联动

[0136_single-number](../0136_single-number)
