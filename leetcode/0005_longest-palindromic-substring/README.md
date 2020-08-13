# 0005_longest-palindromic-substring

遍历数组, 选定一个pivot向两边拓展来判断palindrom.

虽然也有其他思路, 比如dp, 但综合时间和空间, 以及理解的难度, 还是这个方法更直接:

```go
func longestPalindrome(s string) string {
    if s == "" {
      return ""
    }
    maxLen := 1
    runes := []rune(s)
    lenR := len(runes)
    resultRunes := runes[0:1]
    for i, _ := range runes {
      // try odd
      length := 1
      j, k := i-1, i+1
      for j >= 0 && k < lenR {
        if runes[j] != runes[k] {
          break
        }
        length += 2
        j--
        k++
      }
      if length > maxLen {
        maxLen = length
        resultRunes = runes[j+1:k]
      }

      // try even
      length = 0
      j, k = i, i+1
      for j >= 0 && k < lenR {
        if runes[j] != runes[k] {
          break
        }
        length += 2
        j--
        k++
      }
      if length > maxLen {
        maxLen = length
        resultRunes = runes[j+1:k]
      }
    }

    return string(resultRunes)
}
```
