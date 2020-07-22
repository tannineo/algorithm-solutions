# 0055_jump-game

dp.

设`dp[i], 0 <= i < len(nums)`为数字能否从`nums[0]`跳到.

如果`dp[i] = true`, 我们有`dp[i, nums[i]+1)`这个区间都能为`true`.

从左到右遍历即可得解`dp[len(nums) - 1]`.
