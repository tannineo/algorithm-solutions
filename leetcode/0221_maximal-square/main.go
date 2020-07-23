package main

func maximalSquare(matrix [][]byte) int {
	result := 0

	lenI := len(matrix)
	if lenI == 0 {
		return 0
	}
	lenJ := len(matrix[0])

	for i := 0; i < lenI-result; i++ {
		for j := 0; j < lenJ-result; j++ {
			if matrix[i][j] == '1' {
				if result == 0 {
					result = 1
				}

				// start measuring the biggest square from result + 1
				// [i][j] as the left top corner
				localBiggest := measureAt(matrix, i, j, result+1, lenI, lenJ)
				if localBiggest > result {
					result = localBiggest
				}
			}
		}
	}

	return result * result // return the area
}

func measureAt(matrix [][]byte, startI, startJ int, startAs int, limitI, limitJ int) int {
	if startI+startAs > limitI || startJ+startAs > limitJ {
		// we don't have to measure when reach the boundary
		return 0
	}
	// 1. measure startAs X startAs square as base
	for i := startI; i < startI+startAs; i++ {
		for j := startJ; j < startJ+startAs; j++ {
			if matrix[i][j] == '0' {
				return 0
			}
		}
	}

	result := startAs
	// 2. extend the square as much as possible
	for startI+result+1 <= limitI && startJ+result+1 <= limitJ {

		for x := 0; x < result+1; x++ {
			if matrix[startI+x][startJ+result] == '0' {
				return result
			}
			if matrix[startI+x][startJ+result+x] == '0' {
				return result
			}
		}
		result++
	}

	return result
}

func main() {

}
