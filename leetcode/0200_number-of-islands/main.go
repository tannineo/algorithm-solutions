package main

func numIslands(grid [][]byte) int {

	maxI := len(grid)
	if maxI == 0 {
		return 0
	}
	maxJ := len(grid[0])
	if maxJ == 0 {
		return 0
	}
	result := 0

	// 0 <= i < maxI, 0 <= j <= maxJ
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {

				// flood fill the 'island'
				floadfill(grid, i, j, maxI, maxJ)

				result++
			}
		}
	}

	return result
}

func floadfill(grid [][]byte, i, j int, maxI, maxJ int) {
	grid[i][j] = '0'

	// up
	if i-1 >= 0 && grid[i-1][j] == '1' {
		floadfill(grid, i-1, j, maxI, maxJ)
	}
	// left
	if j-1 >= 0 && grid[i][j-1] == '1' {
		floadfill(grid, i, j-1, maxI, maxJ)
	}
	// down
	if i+1 < maxI && grid[i+1][j] == '1' {
		floadfill(grid, i+1, j, maxI, maxJ)
	}
	// right
	if j+1 < maxJ && grid[i][j+1] == '1' {
		floadfill(grid, i, j+1, maxI, maxJ)
	}
}

func main() {
	println(numIslands([][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	})) // 1

	println(numIslands([][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '0', '1'},
	})) // 4
}
