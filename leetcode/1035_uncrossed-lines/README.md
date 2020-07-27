# 1035_uncrossed-lines

注意读题, 每个点只能归属**至多一条线段**, 即不出现一个点连n条线的情况.

dp.

设`dp[i][j]`是序列`A[0, i)`和`B[0, j)`的不相交连线最大值.

- 当`A[i-1] == B[i-1]`时, `dp[i][j] = dp[i-1][j-1] + 1`
- `dp[i][j] = max{ dp[i-1][j], dp[i][j-1] }`

## 联动

[1143_longest-common-subsequence](../1143_longest-common-subsequence)
