package main

func findCircleNum(M [][]int) int {

	// we assume the input is always legal N*N
	N := len(M)
	if N == 0 {
		return 0
	}

	// ensure that the matrix is symmetric
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if M[i][j] == 1 {
				M[j][i] = 1
			}
		}
	}

	// count the 1st line
	circleCount := 0
	for i := 0; i < N; i++ {
		if M[i][i] == 1 {
			circleCount++
			dfsFriend(M, N, i)
		}
	}

	return circleCount
}

// dfsFriend searches all the friends in the circle of `i`
func dfsFriend(M [][]int, N, i int) {
	if M[i][i] == 0 {
		return
	}
	M[i][i] = 0
	for j := 0; j < N; j++ {
		if M[i][j] == 1 {
			dfsFriend(M, N, j)
		}
	}
}
