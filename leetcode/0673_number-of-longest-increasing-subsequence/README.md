# 0673_number-of-longest-increasing-subsequence

dp.

`f(n)`代表: 包含`index=n`位置元素的数组的最长LIS.

`f(n) = max( f(0), ..., f(n-1)) + 1`, `0, ..., n-1`末尾元素必须小于当前元素.

注意需要另开一个数组记录解的个数, 因为一个位置的同一长度可能有复数解.

## 联动

[0300_longest-increasing-subsequence](../0300_longest-increasing-subsequence)
