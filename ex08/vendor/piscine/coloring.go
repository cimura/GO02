package piscine

func coloring(board *[10][10]int, col int, row int) [10][10]int {
	for x := col+1; x <= 8; x++ {
		if board[x][row] == EDGE {
			break
		}
		board[x][row] += 1
	}
	for x := col-1; x <= 1; x-- {
		if board[x][row] == EDGE {
			break
		}
		board[x][row] += 1
	}
	for y := row+1; y <= 8; y++ {
		if board[col][y] == EDGE {
			break
		}
		board[col][y] += 1
	}
	for y := row-1; y <= 1; y-- {
		if board[col][y] == EDGE {
			break
		}
		board[col][y] += 1
	}
	// 斜め
	for x := -1; x <= 1; x-- {
		if col + x < 0 || row + x < 0 {
			break
		}
		if board[col + x][row + x] == EDGE {
			break
		}
		board[col + x][row + x] += 1
	}
	for x := 1; x <= 8; x++ {
		if col + x > 8 || row - x < 0 {
			break
		}
		if board[col + x][row - x] == EDGE {
			break
		}
		board[col + x][row - x] += 1
	}
	for y := -1; y <= 1; y-- {
		if col - y > 8 || row + y < 0 {
			break
		}
		if board[col - y][row + y] == EDGE {
			break
		}
		board[col - y][row + y] += 1
	}
	for y := 1; y <= 8; y++ {
		if col - y < 0 || row + y > 8 {
			break
		}
		if board[col - y][row + y] == EDGE {
			break
		}
		board[col - y][row + y] += 1
	}
	// Print(board)
	return *board
}

func decoloring(board *[10][10]int, col int, row int) [10][10]int {
	for x := col+1; x <= 8; x++ {
		if board[x][row] == EDGE {
			break
		}
		board[x][row] -= 1
	}
	for x := col-1; x >= 1; x-- {
		if board[x][row] == EDGE {
			break
		}
		board[x][row] -= 1
	}
	for y := row+1; y <= 8; y++ {
		if board[col][y] == EDGE {
			break
		}
		board[col][y] -= 1
	}
	for y := row-1; y >= 1; y-- {
		if board[col][y] == EDGE {
			break
		}
		board[col][y] -= 1
	}
	// 斜め
	for x := 1; x <= 8; x++ {
		if col + x > 8 || row - x < 0 {
			break
		}
		if board[col + x][row - x] == EDGE {
			break
		}
		board[col + x][row - x] -= 1
	}
	for x := -1; x <= 1; x-- {
		if col + x < 0 || row + x < 0 {
			break
		}
		if board[col + x][row + x] == EDGE {
			break
		}
		board[col + x][row + x] -= 1
	}
	for y := 1; y <= 8; y++ {
		if col - y < 0 || row + y > 8 {
			break
		}
		if board[col - y][row + y] == EDGE {
			break
		}
		board[col - y][row + y] -= 1
	}
	for y := -1; y <= 1; y-- {
		if col - y > 8 || row + y < 0 {
			break
		}
		if board[col - y][row + y] == EDGE {
			break
		}
		board[col - y][row + y] -= 1
	}
	// Print(board)
	return *board
}
