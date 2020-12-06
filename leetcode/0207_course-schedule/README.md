# 0207_course-schedule

给定课程数目, 和一些前置课程信息, 判断是否可行.

正如题目Example 2给出的, 不能形成回路.

1. 根据prerequisites构建图
   - 构建过程中无法仅靠当前prerequisites pair判断...
   - 但是若出现不属于范围内的课程可以直接返回
2. dfs. 经过的节点做标记. dfs后记得清理现场.

# 联动

[0201_bitwise-and-of-numbers-range](../0201_bitwise-and-of-numbers-range)
