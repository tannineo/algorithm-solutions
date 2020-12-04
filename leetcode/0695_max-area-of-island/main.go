package main

func maxAreaOfIsland(grid [][]int) int {

	maxArea := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				// start search
				area := dfsArea(grid, i, j)
				maxArea = max(area, maxArea)
			}
		}
	}

	return maxArea
}

func dfsArea(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}

	area := 1
	grid[i][j] = 0 // mark as searched

	// up
	if i-1 >= 0 {
		area += dfsArea(grid, i-1, j)
	}
	// down
	if i+1 < len(grid) {
		area += dfsArea(grid, i+1, j)
	}
	// left
	if j-1 >= 0 {
		area += dfsArea(grid, i, j-1)
	}
	// right
	if j+1 < len(grid[i]) {
		area += dfsArea(grid, i, j+1)
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
