package main

func countValidSquaresAsLeftTop(matrix [][]int, i, j int, lenI, lenJ int) int {
	if matrix[i][j] == 0 {
		return 0
	}
	result := 1

	// count from edge = 2
	for edge := 2; i+edge <= lenI && j+edge <= lenJ; edge++ {
		for e := 0; e < edge; e++ {
			if matrix[i+e][j+edge-1] == 0 {
				return result
			}

			if matrix[i+edge-1][j+e] == 0 {
				return result
			}
		}
		result++
	}

	return result
}

func countSquares(matrix [][]int) int {
	result := 0

	// 1 <= i, 1 <= j
	lenI := len(matrix)
	lenJ := len(matrix[0])

	for i := 0; i < lenI; i++ { // row
		for j := 0; j < lenJ; j++ { // col
			result += countValidSquaresAsLeftTop(matrix, i, j, lenI, lenJ)
		}
	}

	return result
}

func main() {

}
