package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 1
	runes := []rune(s)
	lenR := len(runes)
	curRunes := map[rune]int{runes[0]: 0}

	for left, right := 0, 1; right < lenR; right++ {
		if pos, ok := curRunes[runes[right]]; ok {
			for ; left <= pos; left++ {
				delete(curRunes, runes[left])
			}
		}
		curRunes[runes[right]] = right
		result = max(result, right-left+1)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
}
