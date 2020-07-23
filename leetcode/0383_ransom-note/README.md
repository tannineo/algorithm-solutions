# 0383_ransom-note

map.

1. 遍历`magazine`获取字符总数.
2. 遍历`ransomNote`, 每读取一个字符c, `map[c]--`
   - 若`map[c]--`之后为负, 直接return false
3. return true

```go
func canConstruct(ransomNote string, magazine string) bool {
  countMap := make(map[rune]int)

  // 1. count runes in magazine
  for _, v := range magazine {
    countMap[v]++
  }

  // 2. iterate over ransomNote
  for _, v := range ransomNote {
    countMap[v]--
    if countMap[v] < 0 {
      return false
    }
  }

  // 3. pass
  return true
}
```
