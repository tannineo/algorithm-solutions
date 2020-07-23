# 0387_first-unique-character-in-a-string

一个map1维护是否出现, 一个map2维护第一次出现, 扫描完string后遍历map找pos最小值, 或者是遍历string找第一个留在map2里的值(不需要保存位置).

```go
func firstUniqChar(s string) int {
  existMap := make(map[rune]struct{})
  onceMap := make(map[rune]struct{})

  // iterate over s
  for _, v := range s {
    if _, ok := existMap[v]; ok {
      if _, ok := onceMap[v]; ok {
        delete(onceMap, v)
      }
    } else {
      existMap[v] = struct{}{}
      onceMap[v] = struct{}{}
    }
  }

  // find the first element which is in the map
  for i, v := range s {
    if _, ok := onceMap[v]; ok {
      return i
    }
  }

  return -1
}
```

## 联动

如果需要省去最后查找的操作, 需要ordered map或者map -> 双向链表.

[1429_first-unique-number](../1429_first-unique-number)
