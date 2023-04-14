package sudoku

import (
	"regexp"
	"strconv"
)

type Board [9][9]int

func Parse(board string) Board {
	var parsedBoard [9][9]int
	var row [9]int

	board = removeNonDigitsFromString(board)

	if len(board) > 81 {
		panic("please, provide a 9x9 board")
	}

	for i, rune := range board {
		colNumber := i % 9
		if colNumber == 0 {
			row = [9]int{}
		}
		value, _ := strconv.Atoi(string(rune))
		row[colNumber] = value

		rowNumber := i / 9
		parsedBoard[rowNumber] = row
	}

	return parsedBoard
}

func removeNonDigitsFromString(str string) string {
	regex := regexp.MustCompile(`\D`)
	return string(regex.ReplaceAll([]byte(str), []byte("")))
}
