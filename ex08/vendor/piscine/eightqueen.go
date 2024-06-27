package	piscine

import "fmt"

const (
	L = 0
	R = 9
	EDGE = 100
)

func initialize() [10][10]int {
	var board [10][10]int
	for k := 0; k < 10; k++ {
		board[L][k] = EDGE
		board[k][L] = EDGE
		board[R][k] = EDGE
		board[k][R] = EDGE
	}
	return board
}

func recursion(col int, board *[10][10] int, result *string) bool {
	if col > 8 {
		// fmt.Println(result)
		return true
	}
	for row := 1; row <= 8; row++ {
			if board[col][row] == 0 {
				// fmt.Printf("board[%d(i)][%d(row)]\n", i, row)
				*board = coloring(board, col, row)
				// Print(*board)
				if recursion(col + 1, board, result) {
					*result += string(row + '0')
					return true
				}
				*board = decoloring(board, col, row)
			}
		}
		// fmt.Println(array)
		// board = coloring(board, 3, 3)
		return false
	}

func Print(board [10][10]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d	", board[j][i])
		}
		fmt.Println()
	}
}

func EightQueens() {
	board := initialize()
	var result string
	// Print(board)
	if recursion(1, &board, &result) {
		fmt.Println(result)
	}
	// coloring(&board, 2, 2)
	// Print(board)
}