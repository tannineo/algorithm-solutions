# 0525_contiguous-array

直觉是比较经典的"气球涂色"问题.

1. 我们将所有的`0`置换成`-1`
2. 遍历一遍数组, 累加数值形成高低变化的"趋势"
   1. 累加的过程中用一个map记录特定数值的第一次出现位置
   2. 之后若是碰到相同的数值, 则判断该位置与记录的第一次位置的距离, 记录这个值得最大值为解
