package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	leftToRight := make([]int, len(height))
	rightToLeft := make([]int, len(height))

	// left to right
	maxHeight := 0
	for i, v := range height {
		maxHeight = max(maxHeight, v)
		leftToRight[i] = maxHeight
	}

	// right to left
	maxHeight = 0
	for i := len(height) - 1; i >= 0; i-- {
		maxHeight = max(maxHeight, height[i])
		rightToLeft[i] = maxHeight
	}

	// the min of leftToRight and rightToLeft is the surface with water added
	sum := 0
	for i := range height {
		sum += min(leftToRight[i], rightToLeft[i]) - height[i]
	}

	return sum
}

func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
