package main

import (
	"fmt"
	"reflect"
	"sort"
)

func groupAnagramsTLE(strs []string) [][]string {

	result := make([][]string, 0)

	bows := make([]map[rune]int, 0)

	for _, str := range strs {
		// 1. count char BOW
		bowMap := countCharBOW(str)

		// 2. search for existing BOW
		isFound := false
		for i, bow := range bows {
			if reflect.DeepEqual(bow, bowMap) {
				// we find the anagram
				result[i] = append(result[i], str)
				isFound = true
				break
			}
		}

		if !isFound {
			// new kind of anagram
			newStrSlice := []string{str}
			result = append(result, newStrSlice)
			bows = append(bows, bowMap)
		}
	}

	return result
}

func countCharBOW(str string) map[rune]int {
	resultMap := make(map[rune]int)

	for _, runeVal := range str {
		if count, ok := resultMap[runeVal]; ok {
			resultMap[runeVal] = count + 1
		} else {
			resultMap[runeVal] = 1
		}
	}

	return resultMap
}

func groupAnagrams(strs []string) [][]string {

	result := make([][]string, 0)
	savedMap := map[string]int{}

	for _, str := range strs {
		// 1. sort str
		newRunes := []rune(str)
		sort.Slice(newRunes, func(i, j int) bool { return newRunes[i] < newRunes[j] })
		newStr := string(newRunes)

		// 2. search for existing string
		if i, ok := savedMap[newStr]; ok {
			result[i] = append(result[i], str)
		} else {
			result = append(result, []string{str})
			savedMap[newStr] = len(result) - 1
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", groupAnagramsTLE([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	fmt.Printf("%v\n", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
