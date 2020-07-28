# 0072_edit-distance

dp.

可以构建二维dp数组走一遍example, 不难发现规律.

设`dp[i][j]`, `0 <= i <= len(word1)`, `0 <= j <= len(word2)`为`word1[0, i)`和`word2[0, j)`的edit distance.

- 当`word1[i-1] == word2[j-1]`时, `dp[i][j] = dp[i-1][j-1]`.
- 上面条件不成立时, `dp[i][j] = min{ dp[i-1][j], dp[i][j-1], dp[i-1][j-1] } + 1`

从`dp[i-1][j]`或者`dp[i][j-1]`过来的情况就是增加字符.

从`dp[i-1][j-1]`斜向过来的就是替换末尾字符.
