# 0763_partition-labels

水.

说是贪心, 也算模拟. 按照题目的描述计算就行.

两次遍历:

1. 记录字母的最后出现位置.
2. 循环计算, 从开始的地方(一开始是`0`)向后扫, 根据当前字母的最后出现位置更新循环的上限. 返回一个partition后更新开始位置.
