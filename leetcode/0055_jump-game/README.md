# 0055_jump-game

dp.

设`dp[i], 0 <= i < len(nums)`为能否从`nums[0]`跳到`nums[i]`的可能性(true/false).

如果`dp[i] = true`, 我们有`dp[i, i + nums[i] + 1)`这个区间都能为`true`.

从左到右遍历即可得解`dp[len(nums) - 1]`.
