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

func (r Board) Column(number int) [9]int {
	var column [9]int

	for i, iv := range r {
		for j, jv := range iv {
			if j == number {
				column[i] = jv
			}
		}
	}

	return column
}

func (r Board) Row(number int) [9]int {
	var row [9]int

	for i, iv := range r {
		if i == number {
			row = iv
		}
	}

	return row
}

func removeNonDigitsFromString(str string) string {
	regex := regexp.MustCompile(`\D`)
	return string(regex.ReplaceAll([]byte(str), []byte("")))
}
