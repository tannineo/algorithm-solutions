package main

func merge(intervals [][]int) [][]int {
	arr := make([]int, 10001)

	singleInters := map[int]bool{}
	for _, inter := range intervals {
		if inter[0] == inter[1] {
			// special cases
			singleInters[inter[0]] = true
			continue
		}

		arr[inter[0]]++
		arr[inter[1]]--
	}

	ans := [][]int{}
	sum := 0
	var start int
	for i, v := range arr {
		prev := sum
		sum += v

		if prev == 0 && sum > 0 {
			start = i
		} else if prev > 0 && sum == 0 {
			ans = append(ans, []int{start, i})
		} else if prev == 0 && sum == 0 {
			if singleInters[i] {
				ans = append(ans, []int{i, i})
			}
		}
	}

	return ans
}
