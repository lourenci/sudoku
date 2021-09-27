package sudoku_test

import (
	"testing"

	sudoku "github.com/lourenci/sudoku/src"
	"github.com/stretchr/testify/assert"
)

func Test_returns_the_missing_numbers_of_a_cell(t *testing.T) {
	board := sudoku.Board{
		[9]int{6, 0, 0, 0, 2, 3, 0, 0, 0},
		[9]int{7, 4, 0, 1, 0, 0, 0, 0, 9},
		[9]int{0, 0, 1, 4, 0, 0, 6, 2, 8},
		[9]int{0, 9, 5, 6, 0, 0, 0, 0, 0},
		[9]int{0, 1, 0, 2, 3, 0, 0, 6, 0},
		[9]int{2, 0, 0, 5, 0, 0, 8, 4, 0},
		[9]int{1, 0, 6, 0, 0, 0, 4, 5, 0},
		[9]int{0, 0, 0, 0, 4, 1, 0, 8, 0},
		[9]int{0, 7, 3, 8, 5, 0, 0, 0, 2},
	}

	assert.Nil(t, board.MissingNumbersInCell(0, 0))
	assert.Equal(t, []int{4, 7, 8, 9}, board.MissingNumbersInCell(4, 5))
}
