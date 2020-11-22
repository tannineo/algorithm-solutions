package main

import (
	"fmt"
	"sort"
)

// sort.Interface
type WordsToSort []string

func (s WordsToSort) Len() int           { return len(s) }
func (s WordsToSort) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s WordsToSort) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// O(n^2)
func longestStrChain(words []string) int {
	if len(words) == 0 {
		return 0
	}

	// sort by length
	sort.Sort(WordsToSort(words))

	// dp
	dp := make([]int, len(words))

	for i := range words {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if isPredecessor(words[j], words[i]) {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	// find max
	answer := 1
	for _, v := range dp {
		answer = max(answer, v)
	}

	return answer
}

// isPredecessor judge if `strA + ch = strB` (insert a character into strA to be strB)
func isPredecessor(strA, strB string) bool {

	curA, curB := 0, 0
	runesA, runesB := []rune(strA), []rune(strB)

	if len(runesA)+1 != len(runesB) {
		return false
	}

	diffCount := 0
	for curA < len(runesA) && curB < len(runesB) {
		if runesA[curA] == runesB[curB] {
			curA++
			curB++
		} else {
			if diffCount == 0 {
				if curB+1 < len(runesB) && runesA[curA] == runesB[curB+1] {
					curA++
					curB += 2
					diffCount++
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return curA == len(runesA) && (curB == len(runesB) || curB == len(runesB)-1)
}

func main() {
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
}
