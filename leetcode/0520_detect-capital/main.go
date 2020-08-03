package main

import "unicode"

func detectCapitalUse(word string) bool {
	upperCaseCount := 0
	runeWord := []rune(word)
	lenW := len(runeWord)
	for _, c := range runeWord {
		if unicode.IsUpper(c) {
			upperCaseCount++
		}
	}

	if upperCaseCount > 0 {
		if unicode.IsUpper(runeWord[0]) && upperCaseCount == 1 {
			return true
		}
		if upperCaseCount == lenW {
			return true
		}
		return false
	}

	return true // all lower cases
}

func main() {
	println(detectCapitalUse("USA")) // true

	println(detectCapitalUse("leetcode")) // true

	println(detectCapitalUse("Google")) // true

	println(detectCapitalUse("camelCase")) // false
}
