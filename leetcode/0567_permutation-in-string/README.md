# 0567_permutation-in-string

滑动窗口, 统计字符出现次数并比对.

## 联动

[0438_find-all-anagrams-in-a-string](../0438_find-all-anagrams-in-a-string)

拿了这个代码稍作修改

```go
func compareMap(a, b map[rune]int) bool {
	for i := 'a'; i <= 'z'; i++ {
		valA := a[i]
		valB := b[i]
		if valA != valB {
			return false
		}
	}

	return true
}

func checkInclusion(p string, s string) bool {
	lenS := len(s)
	lenP := len(p)

	if lenS < lenP { // p is non-empty
		return false
	}

	charMap := map[rune]int{}

	// init charMap
	for _, c := range p {
		charMap[c]++
	}

	// init window
	runeS := []rune(s)
	windowMap := map[rune]int{}
	for i := 0; i < lenP; i++ {
		windowMap[runeS[i]]++
	}
	// judge init window
	if compareMap(charMap, windowMap) {
		return true
	}
	// slide window
	for i := lenP; i < lenS; i++ {
		windowMap[runeS[i-lenP]]--
		windowMap[runeS[i]]++
		if compareMap(charMap, windowMap) {
			return true
		}
	}

	return false
}
```
