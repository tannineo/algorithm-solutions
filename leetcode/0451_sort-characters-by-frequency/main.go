package main

import "sort"

type RuneCount struct {
	Char  rune
	Count int
}

type ByCount []RuneCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Less(i, j int) bool { return a[i].Count < a[j].Count }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func frequencySort(s string) string {
	lenS := len(s)
	if lenS == 0 {
		return ""
	}

	// 1. count runes
	runeMap := map[rune]int{}
	for _, v := range s {
		runeMap[v]++
	}

	// 2. sort by counts
	arr := []RuneCount{}
	for k, v := range runeMap {
		arr = append(arr, RuneCount{k, v})
	}

	sort.Sort(sort.Reverse(ByCount(arr)))

	answerRunes := []rune{}
	for _, v := range arr {
		for i := 0; i < v.Count; i++ {
			answerRunes = append(answerRunes, v.Char)
		}
	}

	return string(answerRunes)
}

func main() {
	println(frequencySort("tree"))     // eert || eetr
	println(frequencySort("ccaccaaa")) // ccccaaaa
}
