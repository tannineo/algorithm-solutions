# 0202_happy-number

1. 用`map`记录(最直观)
2. Floyd's Cycle-Finding Algorithm:
   - 使用几个不同步长/速度的agent/function去计算, 若是有循环, 必会在某一个时刻求得同一个值
3. 数学结论:
   - 只存在唯一的循环: `4, 16, 37, 58, 89, 145, 42, 20`
