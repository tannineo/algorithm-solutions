package main

import "fmt"

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

func findAnagrams(s string, p string) []int {
	result := []int{}
	lenS := len(s)
	lenP := len(p)

	if lenS < lenP { // p is non-empty
		return result
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
		result = append(result, 0)
	}
	// slide window
	for i := lenP; i < lenS; i++ {
		windowMap[runeS[i-lenP]]--
		windowMap[runeS[i]]++
		if compareMap(charMap, windowMap) {
			result = append(result, i-lenP+1)
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", findAnagrams("cbaebabacd", "abc"))
}
