package main

func maxDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	// check if both `1` and `0` exist
	waterFlag, landFlag := false, false
	for _, v := range grid {
		for _, w := range v {
			if w == 1 {
				landFlag = true
			} else {
				waterFlag = true
			}
		}
	}
	if !(landFlag && waterFlag) {
		return -1
	}

	maxVal := 0
	// for every `0`
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				mapEdges := map[int]bool{}
				mapWater := map[int]bool{}
				// flood fill to add edges
				ffForEdge(grid, i, j, mapEdges, mapWater)
				maxVal = max(maxVal, getMasDistFromMaps(mapEdges, mapWater))
			}
		}
	}

	return maxVal
}

func ffForEdge(grid [][]int, i, j int, mapEdges, mapWater map[int]bool) {
	if grid[i][j] != 0 {
		return
	}

	grid[i][j] = 2                  // mark as visited
	mapWater[coor2Int(i, j)] = true // record to mapWater

	// up
	if i-1 >= 0 {
		if grid[i-1][j] == 1 { // is edge
			mapEdges[coor2Int(i-1, j)] = true
		} else {
			ffForEdge(grid, i-1, j, mapEdges, mapWater)
		}
	}
	// down
	if i+1 < len(grid) {
		if grid[i+1][j] == 1 { // is edge
			mapEdges[coor2Int(i+1, j)] = true
		} else {
			ffForEdge(grid, i+1, j, mapEdges, mapWater)
		}
	}
	// left
	if j-1 >= 0 {
		if grid[i][j-1] == 1 { // is edge
			mapEdges[coor2Int(i, j-1)] = true
		} else {
			ffForEdge(grid, i, j-1, mapEdges, mapWater)
		}
	}
	// right
	if j+1 < len(grid[i]) {
		if grid[i][j+1] == 1 { // is edge
			mapEdges[coor2Int(i, j+1)] = true
		} else {
			ffForEdge(grid, i, j+1, mapEdges, mapWater)
		}
	}
}

func getMasDistFromMaps(mapEdges, mapWater map[int]bool) int {
	maxVal := 0
	for w := range mapWater {
		minVal := 1000 // grid max 100 * 100
		for e := range mapEdges {
			ei, ej := int2Coor(e)
			wi, wj := int2Coor(w)
			minVal = min(minVal, abs(ei-wi)+abs(ej-wj))
		}
		maxVal = max(minVal, maxVal)
	}

	return maxVal
}

func coor2Int(i, j int) int {
	return i*1000 + j
}

func int2Coor(num int) (i, j int) {
	return num / 1000, num % 1000
}

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

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
