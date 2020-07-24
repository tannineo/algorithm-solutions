# 1232_check-if-it-is-a-straight-line

如果`len(coordinates) == 2`, 直接返回`true`.

1. 保存第一个坐标`a`
2. 寻找第一个与`a`不相同的`b`, 若`b`不存在, 或`b`为最后一个元素, 直接返回`true`
3. 通过`a`, `b`计算斜率
4. 遍历余下所有元素, 查看是否符合

注意, 坐标全是int, 小心计算.
