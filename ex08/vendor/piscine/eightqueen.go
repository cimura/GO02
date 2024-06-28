package piscine

const (
	SIZE = 8
	L    = 0
	R    = SIZE + 2
	EDGE = 100
)

func EightQueens() {
	var result string
	board := initialize()
	recursion(1, &board, &result)
}

func recursion(col int, board *[R][R]int, result *string) {
	if col > SIZE {
		printstr(result)
		return
	}
	for row := 1; row <= SIZE; row++ {
		if board[col][row] == 0 {
			*board = coloring(board, col, row)
			Tempresult := *result + string(row + '0')
			recursion(col+1, board, &Tempresult)
			*board = decoloring(board, col, row)
		}
	}
}