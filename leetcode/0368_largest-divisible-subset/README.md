# 0368_largest-divisible-subset

注意, 元素皆唯一.

先从小到大排序, 再dp.

遍历`nums`:

每次`dp[i]++`及之后能被`nums[i]`整除的元素位置dp值`max{当前dp值, 整除的元素位置dp值}`.

找到dp数组里的最大值, 然后反向(dp值依次-1)推出数组.
