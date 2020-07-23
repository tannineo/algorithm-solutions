# 1143_longest-common-subsequence

经典dp.

设`dp[i][j]`是`text1[0, i)`和`text2[0, j)`符合题意的答案(最大公共子序列的长度).

则:

- `dp[0][j] = 0`, `dp[i][0] = 0`
- `dp[i][j] = `:
   1. `dp[i-1][j-1] + 1`: 当`text1[i] == text2[j]`
   2. `max{dp[i-1][j], dp[i][j-1]}`

`dp[i][j]` 为解.

## 联动

[0712_minimum-ascii-delete-sum-for-two-strings](../0712_minimum-ascii-delete-sum-for-two-strings)
