# 0279_perfect-squares

第一反应是dp.

dp当前值是之前所有相差为平方的记录值+1的最小值.

## 最快的数学方法

- 所有非负整数能被 **四个** 整数的二次方 之和表示. 包含0的情况, 所以是1-4
  - https://en.wikipedia.org/wiki/Lagrange%27s_four-square_theorem
- 若一个正整数`n`满足`n = 4^a (8b+7)`, 期中`a`和`b`为非负整数. `n`必为四个非零整数的二次方之和.
  - https://en.wikipedia.org/wiki/Lagrange%27s_four-square_theorem

所以`func numSquares(n int) int`得到答案为`[1, 4]`, 其中答案`1`和`4`的判断都能通过数学方法求得, 而`2`的判断通过遍历, 剩下的就是`3`的情况.
