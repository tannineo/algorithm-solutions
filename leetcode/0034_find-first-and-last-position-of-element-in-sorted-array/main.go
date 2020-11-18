package main

func searchRange(nums []int, target int) []int {
	left := findLeft(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}

	right := findRight(nums, target)

	return []int{left, right}
}

func findLeft(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		cur := (l + r - 1) / 2
		if nums[cur] >= target {
			r = cur
		} else {
			l = cur + 1
		}
	}

	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func findRight(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		cur := (l + r - 1) / 2
		if nums[cur] > target {
			r = cur
		} else {
			l = cur + 1
		}
	}

	l--

	if l < 0 || nums[l] != target {
		return -1
	}
	return l
}
