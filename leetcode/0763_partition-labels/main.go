package main

import "fmt"

func partitionLabels(S string) []int {

	if len(S) == 1 {
		return []int{1}
	}

	// we only care the last appearance of the char
	mapLast := map[byte]int{}
	for i := range S {
		mapLast[S[i]] = i
	}

	ans := []int{}
	start := 0
	for start < len(S) {
		prefix := findPrefixPartiotion(S, start, mapLast)
		ans = append(ans, len(prefix))
		start += len(prefix)
	}

	return ans
}

func findPrefixPartiotion(S string, start int, mapLast map[byte]int) string {
	// len(S) > 0
	last := start
	for i := start; i <= last; i++ {
		// update last
		last = max(last, mapLast[S[i]])
	}

	return S[start : last+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
