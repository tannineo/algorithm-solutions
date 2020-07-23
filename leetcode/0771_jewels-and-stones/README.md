# 0771_jewels-and-stones

map存储:

1. 扫`J`初始化`0`值
2. 扫`S`, 找到map对应的值++
3. 遍历map, 计算总和

```go
func numJewelsInStones(J string, S string) int {

  if S == "" {
    return 0
  }

  countMap := make(map[rune]int)

  // 1. iterate `J` to init 0 values
  for _, v := range J {
    countMap[v] = 0
  }

  // 2. iterate `S` to let values in map +1
  for _, v := range S {
    if _, ok := countMap[v]; ok {
      countMap[v]++
    }
  }

  // 3. iterate map to sum up
  sum := 0
  for _, v := range countMap {
    sum += v
  }

  return sum
}
```
