package main

import (
	"fmt"
	"sort"
	"strings"
)

var MapIndex map[string]int

type LogsToSort []string

func (l LogsToSort) Len() int      { return len(l) }
func (l LogsToSort) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l LogsToSort) Less(i, j int) bool {
	if isDigitLog(l[i]) && isDigitLog(l[j]) {
		// both are digit-logs, sort by index
		return MapIndex[l[i]] < MapIndex[l[j]]
	} else if !(isDigitLog(l[i]) || isDigitLog(l[j])) {
		// both are letter-logs, sort lexicographically
		return isLessLetterLog(l[i], l[j])
	} else {
		// letter-logs < digit-logs
		return !isDigitLog(l[i])
	}
}

func isDigitLog(log string) bool {
	strs := strings.Split(log, " ")
	return strs[1][0] >= '0' && strs[1][0] <= '9'
}

func isLessLetterLog(logI, logJ string) bool {
	idI, wordsI := log2words(logI)
	idJ, wordsJ := log2words(logJ)

	if wordsI == wordsJ {
		return idI < idJ
	}
	return wordsI < wordsJ
}

func log2words(log string) (identifier, words string) {
	seps := strings.Split(log, " ")
	return seps[0], strings.Join(seps[1:], " ")
}

func reorderLogFiles(logs []string) []string {
	MapIndex = map[string]int{}
	for i, log := range logs {
		MapIndex[log] = i
	}
	sort.Sort(LogsToSort(logs))
	return logs
}

func main() {
	ans := reorderLogFiles([]string{
		"8 fj dzz k",
		"5r 446 6 3",
		"63 gu psub",
		"5 ba iedrz",
		"6s 87979 5",
		"3r 587 01",
		"jc 3480612",
		"bb wsrd kp",
		"b aq cojj",
		"r5 6316 71",
	})

	fmt.Println(ans)
}
