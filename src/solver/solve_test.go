package solver_test

import (
	"testing"

	sudoku "github.com/lourenci/sudoku/src"
	"github.com/lourenci/sudoku/src/solver"
	"github.com/stretchr/testify/assert"
)

func Test_solves_the_board(t *testing.T) {
	board := sudoku.Board{
		{0, 4, 9, 0, 0, 0, 0, 3, 0},
		{0, 5, 0, 6, 1, 0, 0, 0, 0},
		{0, 0, 8, 0, 2, 9, 5, 0, 6},
		{8, 0, 0, 9, 0, 7, 0, 0, 4},
		{7, 0, 0, 0, 0, 0, 0, 8, 1},
		{0, 2, 5, 0, 4, 1, 3, 0, 0},
		{2, 0, 0, 0, 7, 6, 0, 1, 0},
		{5, 0, 0, 4, 0, 8, 7, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 9, 5},
	}

	expectedBoard := sudoku.Board{
		{6, 4, 9, 7, 8, 5, 1, 3, 2},
		{3, 5, 2, 6, 1, 4, 8, 7, 9},
		{1, 7, 8, 3, 2, 9, 5, 4, 6},
		{8, 3, 1, 9, 6, 7, 2, 5, 4},
		{7, 6, 4, 2, 5, 3, 9, 8, 1},
		{9, 2, 5, 8, 4, 1, 3, 6, 7},
		{2, 9, 3, 5, 7, 6, 4, 1, 8},
		{5, 1, 6, 4, 9, 8, 7, 2, 3},
		{4, 8, 7, 1, 3, 2, 6, 9, 5},
	}

	assert.Equal(t, expectedBoard, solver.Solve(board))
}
