# 0344_reverse-string

水爆.

```go
func reverseString(s []byte)  {
  i := 0
  j := len(s)-1
  for ; i < j;  {
    s[i], s[j] = s[j], s[i]
    i++
    j--
  }
}
```
