# 0091_decode-ways

分治, 递归时作记忆优化, 记录一些重复出现的字符串的答案, 否则会TLE.

将字符串从中间分开递归, 若分开的地方, 左右两个数字能组成一个`A-Z`, 则去掉这两个数字, 再一次对左右递归.
