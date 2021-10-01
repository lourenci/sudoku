package parsers

import (
	"strconv"

	"github.com/lourenci/sudoku"
)

type StringBoard struct {}

func (s StringBoard) Parse(input string) sudoku.Board {
	var b sudoku.Board

	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			number, _ := strconv.Atoi(string(input[i*9+j]))
			b[i][j] = number
		}
	}

	return b
}
