package sudoku_test

import (
	"testing"

	"lourenci.com/sudoku/modules/assert"
	"lourenci.com/sudoku/sudoku"
)

func TestParse(t *testing.T) {
	assert.Equal(
		t,
		sudoku.Parse("12345678923"),
		sudoku.Board(
			[9][9]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{2, 3, 0, 0, 0, 0, 0, 0, 0},
			},
		),
	)
}
