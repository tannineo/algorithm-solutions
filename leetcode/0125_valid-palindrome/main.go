package main

import "unicode"

func isPalindrome(s string) bool {
	// filtered letters, storing lower cases
	filteredRunes := []rune{}
	for _, r := range s {
		if unicode.IsUpper(r) {
			filteredRunes = append(filteredRunes, unicode.ToLower(r))
		} else if unicode.IsLower(r) {
			filteredRunes = append(filteredRunes, r)
		}
	}

	lenR := len(filteredRunes)

	// judge palindrome
	for i, j := 0, lenR-1; i < j; {
		if filteredRunes[i] != filteredRunes[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	println(isPalindrome("A man, a plan, a canal: Panama")) // true

	println(isPalindrome("race a car")) // false
}
