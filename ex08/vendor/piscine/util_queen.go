package piscine

import "ft"

func initialize() [R][R]int {
	var board [R][R]int
	for k := L; k < R; k++ {
		board[L][k] = EDGE
		board[k][L] = EDGE
		board[R-1][k] = EDGE
		board[k][R-1] = EDGE
	}
	return board
}

func coloring(board *[R][R]int, col int, row int) [R][R]int {
	new_board := *board
	new_board[col][row] = 1
	for x := col + 1; x <= SIZE; x++ {
		if new_board[x][row] == EDGE {
			break
		}
		new_board[x][row] += 1
	}
	for x := col - 1; x <= SIZE; x-- {
		if new_board[x][row] == EDGE {
			break
		}
		new_board[x][row] += 1
	}
	for y := row + 1; y <= SIZE; y++ {
		if new_board[col][y] == EDGE {
			break
		}
		new_board[col][y] += 1
	}
	for y := row - 1; y <= SIZE; y-- {
		if new_board[col][y] == EDGE {
			break
		}
		new_board[col][y] += 1
	}
	// 斜め
	// 右下
	for i := 1; col+i <= SIZE && row+i <= SIZE; i++ {
		if new_board[col+i][row+i] == EDGE {
			break
		}
		new_board[col+i][row+i] += 1
	}
	// 右上
	for i := 1; col+i <= SIZE && row-i >= 0; i++ {
		if new_board[col+i][row-i] == EDGE {
			break
		}
		new_board[col+i][row-i] += 1
	}
	// 左下
	for i := 1; col-i >= 0 && row+i <= SIZE; i++ {
		if new_board[col-i][row+i] == EDGE {
			break
		}
		new_board[col-i][row+i] += 1
	}
	// 左上
	for i := 1; col-i >= 0 && row-i >= 0; i++ {
		if new_board[col-i][row-i] == EDGE {
			break
		}
		new_board[col-i][row-i] += 1
	}
	// Print(new_board)
	return new_board
}

func decoloring(board *[R][R]int, col int, row int) [R][R]int {
	new_board := *board
	new_board[col][row] = 0
	for x := col + 1; x <= SIZE; x++ {
		if new_board[x][row] == EDGE {
			break
		}
		new_board[x][row] -= 1
	}
	for x := col - 1; x <= SIZE; x-- {
		if new_board[x][row] == EDGE {
			break
		}
		new_board[x][row] -= 1
	}
	for y := row + 1; y <= SIZE; y++ {
		if new_board[col][y] == EDGE {
			break
		}
		new_board[col][y] -= 1
	}
	for y := row - 1; y <= SIZE; y-- {
		if new_board[col][y] == EDGE {
			break
		}
		new_board[col][y] -= 1
	}
	// 斜め
	// 右下
	for i := 1; col+i <= SIZE && row+i <= SIZE; i++ {
		if new_board[col+i][row+i] == EDGE {
			break
		}
		new_board[col+i][row+i] -= 1
	}
	// 右上
	for i := 1; col+i <= SIZE && row-i >= 0; i++ {
		if new_board[col+i][row-i] == EDGE {
			break
		}
		new_board[col+i][row-i] -= 1
	}
	// 左下
	for i := 1; col-i >= 0 && row+i <= SIZE; i++ {
		if new_board[col-i][row+i] == EDGE {
			break
		}
		new_board[col-i][row+i] -= 1
	}
	// 左上
	for i := 1; col-i >= 0 && row-i >= 0; i++ {
		if new_board[col-i][row-i] == EDGE {
			break
		}
		new_board[col-i][row-i] -= 1
	}
	// Print(new_board)
	return new_board
}

func printstr(result *string) {
	for _, r := range *result {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}