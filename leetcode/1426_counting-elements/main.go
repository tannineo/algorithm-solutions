package main

func countElements(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	countMap := make(map[int]int)
	result := 0

	// 1. iterate over the array to count the numbers
	for _, v := range arr {
		countMap[v]++
	}

	// 2. iterate over the map to find the numbers that meet the condition
	for k, cnt := range countMap {
		if _, ok := countMap[k+1]; ok {
			result += cnt
		}
	}

	return result
}

func main() {
	println(countElements([]int{1, 2, 3})) // 2

	println(countElements([]int{1, 1, 3, 3, 5, 5, 7, 7})) // 0

	println(countElements([]int{1, 3, 2, 3, 5, 0})) // 3

}
