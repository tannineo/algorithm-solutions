package main

func longestSubsequence(arr []int, difference int) int {
	storeMap := map[int]int{}

	for _, v := range arr {
		last := v - difference

		if len, ok := storeMap[last]; ok {
			storeMap[v] = len + 1
		} else {
			storeMap[v] = 1
		}
	}

	maxLen := 0
	for _, v := range storeMap {
		if maxLen < v {
			maxLen = v
		}
	}

	return maxLen
}
