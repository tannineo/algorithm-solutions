package main

func solve(board [][]byte) {
	m := len(board)
	if m <= 1 {
		return
	}

	n := len(board[0])
	if n <= 1 {
		return
	}

	// 1. fill edge `O` with `E`
	for i := 0; i < m; i++ {
		fillWith(board, m, n, i, 0, 'O', 'E')
		fillWith(board, m, n, i, n-1, 'O', 'E')
	}
	for j := 0; j < n; j++ {
		fillWith(board, m, n, 0, j, 'O', 'E')
		fillWith(board, m, n, m-1, j, 'O', 'E')
	}

	// 2. replace rest `O` with `X`
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	// 3. replace `E` back to `O`
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'E' {
				board[i][j] = 'O'
			}
		}
	}
}

func fillWith(board [][]byte, m, n, i, j int, is byte, with byte) {
	if board[i][j] != is {
		return
	}

	board[i][j] = with

	if i-1 >= 0 {
		fillWith(board, m, n, i-1, j, is, with)
	}
	if i+1 < m {
		fillWith(board, m, n, i+1, j, is, with)
	}
	if j-1 >= 0 {
		fillWith(board, m, n, i, j-1, is, with)
	}
	if j+1 < n {
		fillWith(board, m, n, i, j+1, is, with)
	}
}

func main() {}
