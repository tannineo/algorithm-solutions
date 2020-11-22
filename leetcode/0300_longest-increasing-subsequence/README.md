# 0300_longest-increasing-subsequence

DP.

`O(n^2)`:

从左到右计算 以每个元素为结尾的序列的最长递增长度. 当前元素的答案是之前所有答案中最长, 且结尾元素小于当前值.
最后遍历获得一个最长长度.

`O(n logn)`:

动态地构建最长递增子序列.

从左到右遍历, 将元素按大小顺序插入正在构建的dp数组, 插入要求是该位置上的元素是**第一个不小于它的数字**, 或者最后:

如果插入的元素在末尾, 则序列长度`+1`.

插入时用Binary Search优化插入时间. 注意dp数组**并不是**LIS

详见Solution: https://leetcode.com/problems/longest-increasing-subsequence/solution/
