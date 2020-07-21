package main

func search(nums []int, target int) int {
	return helper3(nums, 0, len(nums), target)
}

// helper3 takes the array `nums` and two int `a` `b`
// to find the `target`
// 0 <= a <= b <= len(Nums)
// search area [a, b)
func helper3(nums []int, a, b int, target int) int {
	lenAB := b - a
	// short enough, directly iterate
	if lenAB <= 4 {
		for i := a; i < b; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	// consider the range	[a, p, q, b)
	p := a + lenAB/3
	q := p + lenAB/3

	// check a p q b
	switch target {
	case nums[a]:
		return a
	case nums[p]:
		return p
	case nums[q-1]:
		return q - 1
	case nums[b-1]:
		return b - 1
	}

	if nums[p] > nums[q-1] {
		// the pivot is in [p, q)
		if target < nums[p] && target > nums[q-1] {
			// binary search
			left := helper2(nums, a, p, target)
			right := helper2(nums, q, b, target)
			if left != -1 {
				return left
			}
			if right != -1 {
				return right
			}
			return -1
		}
		return helper3(nums, p, q, target)
	} else {
		// the pivot is in [0, p) or [q, len)
		if nums[p] < target && target < nums[q-1] {
			return helper3(nums, p, q, target)
		}

		left := helper3(nums, a, p, target)
		right := helper3(nums, q, b, target)
		if left != -1 {
			return left
		}
		if right != -1 {
			return right
		}
		return -1
	}
}

func helper2(nums []int, a, b int, target int) int {
	lenAB := b - a
	// short enough, directly iterate
	if lenAB <= 2 {
		for i := a; i < b; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	// binary search
	p := (a + b) / 2
	// check a p b
	switch target {
	case nums[a]:
		return a
	case nums[p]:
		return p
	case nums[b-1]:
		return b - 1
	}

	// consider [a, p, b)
	if target < nums[p] {
		return helper2(nums, a, p, target)
	} else {
		return helper2(nums, p+1, b, target)
	}
}

func main() {
	// nums[3] = 7, pivot = 3
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4

	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1

	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 1)) // 5

	println(search([]int{6, 7, 8, 1, 2, 3, 4, 5}, 3)) // 5
}
