package parsers_test

import (
	"testing"

	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/parsers"
	"github.com/stretchr/testify/assert"
)

func Test_parses_a_string_into_a_board(t *testing.T) {
	stringBoard := "094050700035082004001000050002005000400193006000000370000000200563824197009070403"

	expectedBoard := sudoku.Board{
		{0, 9, 4, 0, 5, 0, 7, 0, 0},
		{0, 3, 5, 0, 8, 2, 0, 0, 4},
		{0, 0, 1, 0, 0, 0, 0, 5, 0},
		{0, 0, 2, 0, 0, 5, 0, 0, 0},
		{4, 0, 0, 1, 9, 3, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 3, 7, 0},
		{0, 0, 0, 0, 0, 0, 2, 0, 0},
		{5, 6, 3, 8, 2, 4, 1, 9, 7},
		{0, 0, 9, 0, 7, 0, 4, 0, 3},
	}

	assert.Equal(t, expectedBoard, parsers.StringBoard{}.Parse(stringBoard))
}
