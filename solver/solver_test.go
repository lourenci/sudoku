package solver_test

import (
	"testing"

	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
	"github.com/lourenci/sudoku/solver/strategies"
	"github.com/stretchr/testify/assert"
)

func Test_solves_the_board_using_the_specified_strategies(t *testing.T) {
	board := sudoku.Board{
		{0, 9, 4, 0, 5, 0, 7, 0, 0},
		{0, 3, 5, 0, 8, 2, 0, 0, 4},
		{0, 0, 1, 0, 0, 0, 0, 5, 0},
		{0, 0, 2, 0, 0, 5, 0, 0, 0},
		{4, 0, 0, 1, 0, 3, 0, 0, 6},
		{0, 0, 0, 0, 0, 0, 3, 7, 0},
		{0, 0, 0, 0, 0, 0, 2, 0, 0},
		{5, 6, 0, 8, 2, 0, 0, 9, 7},
		{0, 0, 9, 0, 7, 0, 4, 0, 3},
	}

	expectedBoard := sudoku.Board{
		{2, 9, 4, 6, 5, 1, 7, 3, 8},
		{7, 3, 5, 9, 8, 2, 6, 1, 4},
		{6, 8, 1, 4, 3, 7, 9, 5, 2},
		{3, 1, 2, 7, 6, 5, 8, 4, 9},
		{4, 7, 8, 1, 9, 3, 5, 2, 6},
		{9, 5, 6, 2, 4, 8, 3, 7, 1},
		{8, 4, 7, 3, 1, 9, 2, 6, 5},
		{5, 6, 3, 8, 2, 4, 1, 9, 7},
		{1, 2, 9, 5, 7, 6, 4, 8, 3},
	}

	assert.Equal(
		t,
		expectedBoard,
		solver.NewSolver([]solver.Strategy{strategies.HiddenSingle{}, strategies.NakedSingle{}}).Solve(board),
	)
}
