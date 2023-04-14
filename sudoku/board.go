package sudoku

import (
	"strconv"
)

type Board [9][9]int

func Parse(board string) Board {
	var parsedBoard [9][9]int
	var row [9]int

	for i, rune := range board {
		colNumber := i % 9
		if colNumber == 0 {
			row = [9]int{}
		}

		value, _ := strconv.Atoi(string(rune))
		row[colNumber] = value

		if i > 0 && (i%8 == 0 || i+1 == len(board)) {
			rowNumber := i / 9
			parsedBoard[rowNumber] = row
		}
	}

	return parsedBoard
}
