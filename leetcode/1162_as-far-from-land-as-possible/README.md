# 1162_as-far-from-land-as-possible

遍历所有`0`格:

1. flood fill相邻的`0`(标记为`2`), 找到所有边缘的`1`, 将水面和边缘分别记录到两个set(map).
2. 单次flood fill之后, 遍历set中的边缘和水面以计算曼哈顿距离, 求最小值.
