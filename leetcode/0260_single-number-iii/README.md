# 0260_single-number-iii

之前类似的题目是找一个数. 现在是两个

1. 遍历`nums`做`XOR`, 得到两个single number的异或值
2. 异或值其中的某一个位置`1`代表了这两个数, 在这个位置上一个为`1`, 一个为`0`
3. 第二次遍历时, 遍历异或, 但是有过滤条件: 这个位置得有`1`
   - 因为其他数字都出现两次, 即使通过过滤也不会对结果造成影响
4. 第二次的异或值就是其中一个数字, `XOR`第一次的异或值就是另一个数字

```go
func singleNumber(nums []int) []int {
    xorResult := 0
    for _, num := range nums {
        xorResult ^= num
    }

    // get the right-most difference
    diff := xorResult & (-xorResult)

    x := 0
    for _, num := range nums {
        if num & diff != 0 {
            x ^= num
        }
    }

    return []int{x, x^xorResult}
}
```

## 联动

[0136_single-number](./single-number)
