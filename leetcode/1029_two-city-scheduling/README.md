# 1029_two-city-scheduling

假定所有人都去A, 根据每个人到A城和B城的差值排序.

取去B城减少费用最多(或者增加费用最小)的一半去B城.

再苛刻点, 就是做partial sort: 手写快排附加判断.
