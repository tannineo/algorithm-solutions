# 0827_making-a-large-island

思路挺明确的, 还是dfs/flood fill.

几次遍历:

1. 计算每个岛屿的面积, 然后给岛屿每一块土地一个负数编号, 把岛屿的面积存到一个map: 编号 -> 面积.
2. 遍历`0`格. 查询计算其上下左右岛屿面积包括自身的和. (有`0`不存在的情况, 即全是陆地)
