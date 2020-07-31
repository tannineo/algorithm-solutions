# 0518_coin-change-2

dp.

`dp[i]`是`i`的找零总方法数. `dp[0] = 1`: 不用coin我们就能凑出0.

循环`j`次, `j`是coin数目, 每次从`coin[j]`开始:

- `dp[i] = dp[i - coin[j]] + dp[i]`

`dp[amount]`就是解
