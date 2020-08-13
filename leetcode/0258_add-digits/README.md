# 0258_add-digits

题目的要求提示我们使用数学方法.

其实多列些数字找规律就能发现.

```go
func addDigits(num int) int {
    if num == 0 {
      return 0
    }
    if num % 9 == 0 {
      return 9
    }
    return num % 9
}
```

详细的数学证明见[solution](https://leetcode.com/problems/add-digits/solution/)
