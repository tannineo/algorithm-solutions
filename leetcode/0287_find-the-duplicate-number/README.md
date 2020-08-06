# 0287_find-the-duplicate-number

题目设下了很多限制.

看解法关于龟兔探测循环算法的解释更快:

https://leetcode.com/problems/find-the-duplicate-number/solution/

```go
func findDuplicate(nums []int) int {
  // 1. 龟兔
  hare := nums[0]
  tortoise := nums[0]
  for {
    // 2. 龟一步 兔两步
    tortoise = nums[tortoise]
    hare = nums[nums[hare]]
    if tortoise == hare {
      break // 落到同一个数字时break
    }
  }

  // 3. 乌龟从起点开始
  // 兔速度变为1
  // 注意一开始就是答案的情况
  tortoise = nums[0]
  for hare != tortoise { // 再次相遇的地方为答案
    tortoise = nums[tortoise]
    hare = nums[hare]
  }

  return tortoise
}
```
