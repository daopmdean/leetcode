package medium

func isValidSudoku(board [][]byte) bool {
	// row validation
	for i := 0; i < 9; i++ {
		m := map[byte]bool{}
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if m[board[i][j]] {
				return false
			}
			m[board[i][j]] = true
		}
	}

	// column validation
	for i := 0; i < 9; i++ {
		m := map[byte]bool{}
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if m[board[j][i]] {
				return false
			}
			m[board[j][i]] = true
		}
	}

	// box validation
	for i := 0; i < 9; i++ {
		m := map[byte]bool{}
		r, c := (i/3)*3, (i%3)*3
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if board[r+j][c+k] == '.' {
					continue
				}
				if m[board[r+j][c+k]] {
					return false
				}
				m[board[r+j][c+k]] = true
			}
		}
	}

	return true
}
