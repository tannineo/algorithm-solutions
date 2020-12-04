package main

func largestIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	maxArea := 0
	mapAreas := map[int]int{}
	count := 2 // -1 is used
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				// start search
				area := dfsArea(grid, i, j)
				mapAreas[-count] = area
				// fill the area
				dfsFillArea(grid, -count, i, j)
				count++
			}
		}
	}

	// find 0 and the sum areas nearby
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				area := 1
				mapAdded := map[int]int{}
				// up
				if i-1 >= 0 && grid[i-1][j] < -1 {
					mapAdded[grid[i-1][j]] = mapAreas[grid[i-1][j]]
				}
				// down
				if i+1 < len(grid) && grid[i+1][j] < -1 {
					mapAdded[grid[i+1][j]] = mapAreas[grid[i+1][j]]
				}
				// left
				if j-1 >= 0 && grid[i][j-1] < -1 {
					mapAdded[grid[i][j-1]] = mapAreas[grid[i][j-1]]
				}
				// right
				if j+1 < len(grid[i]) && grid[i][j+1] < -1 {
					mapAdded[grid[i][j+1]] = mapAreas[grid[i][j+1]]
				}
				for _, v := range mapAdded {
					area += v
				}
				maxArea = max(maxArea, area)
			}
		}
	}

	if maxArea == 0 { // all 1 in the grid
		maxArea = len(grid) * len(grid[0])
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfsArea(grid [][]int, i, j int) int {
	if grid[i][j] != 1 {
		return 0
	}

	area := 1
	grid[i][j] = -1 // mark as searched

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

func dfsFillArea(grid [][]int, area, i, j int) {
	if grid[i][j] != -1 {
		return
	}

	grid[i][j] = area // fill the area

	// up
	if i-1 >= 0 {
		dfsFillArea(grid, area, i-1, j)
	}
	// down
	if i+1 < len(grid) {
		dfsFillArea(grid, area, i+1, j)
	}
	// left
	if j-1 >= 0 {
		dfsFillArea(grid, area, i, j-1)
	}
	// right
	if j+1 < len(grid[i]) {
		dfsFillArea(grid, area, i, j+1)
	}
}
