# 0303_range-sum-query-immutable

先行计算.

`sumRange(i, j) = sum - sumRange(0, i) - sumRange(j, len(arr))`

上述范围都是左闭右开(`[i, j)`), 所以注意边界处理.
