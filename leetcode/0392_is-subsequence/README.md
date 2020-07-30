# 0392_is-subsequence

两个cursor遍历判断即可

```go
func isSubsequence(s string, t string) bool {
  if s == "" {
    return true
  }

  if t == "" {
    return false
  }

  curS := 0
  curT := 0
  lenS := len(s)
  lenT := len(t)
  runesS := []rune(s)
  runesT := []rune(t)
  for {
    if runesS[curS] == runesT[curT] {
      curS++
    }
    curT++
    if curT >= lenT && curS < lenS {
      return false
    }
    if curS >= lenS {
      return true
    }
  }

  return true
}
```
