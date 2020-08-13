package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	maxLen := 1
	runes := []rune(s)
	lenR := len(runes)
	resultRunes := runes[0:1]

	for i := range runes {
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
			resultRunes = runes[j+1 : k]
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
			resultRunes = runes[j+1 : k]
		}
	}

	return string(resultRunes)
}

func main() {
	println(longestPalindrome("ccc"))
}
