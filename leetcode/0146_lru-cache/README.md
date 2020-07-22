# 0146_lru-cache

Least Recently Used (LRU) cache.

要求使用O(1)的时间完成存取操作.

单纯的存取用map完成. 此处我们需要维护一个双向链表.

map存着链表的node, 双向链表使得链表摘取单个node到尾部操作变为O(1).
