# 0997_find-the-town-judge

题目条件:

1. town judge 不信赖任何人
2. town judge 被所有人信赖
3. 只有一个人符合town judge的条件

遍历`trust`, 维护两个map

- map1 记录对他人采取了信赖的人的编号
- map2 记录被别人信赖, (且未信赖他人的, 如果进行删除操作), town judge候补的被信赖次数

遍历结束后, 遍历map2, 若被信赖次数达到`N - 1`且这个数唯一, 则是town judge.
