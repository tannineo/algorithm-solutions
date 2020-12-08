package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {

	// shrink
	// all the numbers remain atmost 3
	mapInsCount := map[int]int{}
	for _, num := range nums {
		if mapInsCount[num] < 3 {
			mapInsCount[num]++
		}
	}

	nums = []int{}
	for num, count := range mapInsCount {
		for i := 0; i < count; i++ {
			nums = append(nums, num)
		}
	}

	mapIntToArr := map[int][][2]int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sumOf2 := nums[i] + nums[j]
			if _, ok := mapIntToArr[sumOf2]; !ok {
				mapIntToArr[sumOf2] = [][2]int{}
			}
			mapIntToArr[sumOf2] = append(mapIntToArr[sumOf2], [2]int{i, j})
		}
	}

	mapSetRecord := map[string]bool{}
	mapSearched := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if mapSearched[nums[i]] {
			continue
		}
		mapSearched[nums[i]] = true
		if _, ok := mapIntToArr[-nums[i]]; ok {
			for _, ids := range mapIntToArr[-nums[i]] {
				if i != ids[0] && i != ids[1] {
					mapSetRecord[sort3NumsToString(nums[i], nums[ids[0]], nums[ids[1]])] = true
				}
			}
		}
	}

	ans := [][]int{}
	for k := range mapSetRecord {
		ans = append(ans, stringTo3Nums(k))
	}

	return ans
}

func sort3NumsToString(i, j, k int) string {
	toSort := []int{i, j, k}
	sort.Ints(toSort)

	return fmt.Sprintf("%v %v %v", toSort[0], toSort[1], toSort[2])
}

func stringTo3Nums(str string) []int {
	strs := strings.Split(str, " ")

	ans := []int{}
	for _, s := range strs {
		s, err := strconv.Atoi(s)
		if err == nil {
			ans = append(ans, s)
		}
	}
	return ans
}
