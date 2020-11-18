package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		cur := (r-l)/2 + l

		if nums[cur] == target {
			return cur
		} else if nums[cur] < target {
			l = cur + 1
		} else {
			r = cur
		}
	}

	return -1
}
