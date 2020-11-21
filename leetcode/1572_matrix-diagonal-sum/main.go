package main

func diagonalSum(mat [][]int) int {
	if len(mat) == 1 {
		return mat[0][0]
	}

	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i] + mat[len(mat)-1-i][i]
	}

	if len(mat)%2 == 1 {
		sum -= mat[len(mat)/2][len(mat)/2]
	}

	return sum
}
