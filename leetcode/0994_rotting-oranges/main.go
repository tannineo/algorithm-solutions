package main

func orangesRotting(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	// start rotting
	answer := 0
	for {
		stillRotting := false
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == 1 {
					stillRotting = tryRotting(grid, m, n, i, j, answer+2) || stillRotting
				}
			}
		}
		if !stillRotting {
			break
		}
		answer++
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 { // fresh oranges left
				return -1
			}
		}
	}

	return answer
}

func tryRotting(grid [][]int, m, n, i, j, currentPhase int) bool {
	if i-1 >= 0 && grid[i-1][j] == currentPhase {
		grid[i][j] = currentPhase + 1
		return true
	}
	if i+1 < m && grid[i+1][j] == currentPhase {
		grid[i][j] = currentPhase + 1
		return true
	}
	if j-1 >= 0 && grid[i][j-1] == currentPhase {
		grid[i][j] = currentPhase + 1
		return true
	}
	if j+1 < n && grid[i][j+1] == currentPhase {
		grid[i][j] = currentPhase + 1
		return true
	}
	return false
}

func main() {
	println(orangesRotting([][]int{
		{2, 1, 1}, {0, 1, 1}, {1, 0, 1}, // -1
	}))

	println(orangesRotting([][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1}, // 4
	}))
}
