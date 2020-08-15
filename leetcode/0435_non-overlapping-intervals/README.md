# 0435_non-overlapping-intervals

DP, greedy.

我们现将`intervals`以end从小到大排序.

遍历`intervals`, 维护当前能保留的interval数量以及整个保留着的interval的最右侧.

当前interval的start若是大于等于最右侧则加入.

答案是总数减去保留的interval数量.
