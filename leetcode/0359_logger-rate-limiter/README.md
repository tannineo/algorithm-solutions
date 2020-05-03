# 0359_logger-rate-limiter

https://leetcode.com/problems/logger-rate-limiter/

用map去存log的记录

值得注意的前提:

1. 过来的log应该是按照实际情况增大timestamp的

实际使用过程中:

1. log的timestamp不一定完全递增
2. 用来记录的map需要定时(?)清理
