package sudoku

import (
	"strconv"
)

type Board [9][9]int

func Parse(board string) Board {
	var parsedBoard [9][9]int
	var row [9]int

	for i, rune := range board {
		if i%9 == 0 {
			row = [9]int{}
		}

		value, _ := strconv.Atoi(string(rune))
		cellNumber := i % 9
		row[cellNumber] = value

		if i > 0 && (i%8 == 0 || i+1 == len(board)) {
			rowNumber := i / 9
			parsedBoard[rowNumber] = row
		}
	}

	return parsedBoard
}
