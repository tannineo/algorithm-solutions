# 0918_maximum-sum-circular-subarray

## dp

第一反应是dp.

**在不考虑循环的情况**: 计算`dp[i]`对于`A[0, i)`的max sum suffix.

`dp[i] = max{ dp[i-1], 0 } + A[i-1]`

`max(dp)`即答案

**在考虑循环的情况下**, 需要记录所得`dp[len(A)]`suffix的起始`start`.

计算`A[0, lenA)`的max sum preffix`maxSumPreffix`.

`dp[len(A)]`所记录的start会在计算过程中被右移.

## 优化

dp数组可以简化, 因为我们只用到了上次一的值.
