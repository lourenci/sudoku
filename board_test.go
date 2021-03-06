package sudoku_test

import (
	"testing"

	"github.com/lourenci/sudoku"
	"github.com/stretchr/testify/assert"
)

func Test_returns_the_missing_numbers_of_a_cell(t *testing.T) {
	board := sudoku.Board{
		{6, 0, 0, 0, 2, 3, 0, 0, 0},
		{7, 4, 0, 1, 0, 0, 0, 0, 9},
		{0, 0, 1, 4, 0, 0, 6, 2, 8},
		{0, 9, 5, 6, 0, 0, 0, 0, 0},
		{0, 1, 0, 2, 3, 0, 0, 6, 0},
		{2, 0, 0, 5, 0, 0, 8, 4, 0},
		{1, 0, 6, 0, 0, 0, 4, 5, 0},
		{0, 0, 0, 0, 4, 1, 0, 8, 0},
		{0, 7, 3, 8, 5, 0, 0, 0, 2},
	}

	assert.Nil(t, board.MissingNumbersInCell(sudoku.Coordinate{0, 0}))
	assert.Equal(t, []int{4, 7, 8, 9}, board.MissingNumbersInCell(sudoku.Coordinate{4, 5}))
	assert.Equal(t, []int{7, 9}, board.MissingNumbersInCell(sudoku.Coordinate{0, 3}))
}

func Test_returns_if_board_is_complete(t *testing.T) {
	incompleteBoard := sudoku.Board{
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

	assert.Equal(t, incompleteBoard.IsComplete(), false)

	completeBoard := sudoku.Board{
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

	assert.Equal(t, completeBoard.IsComplete(), true)
}

func Test_gets_the_start_and_end_coordinates_of_a_give_house(t *testing.T) {
	startCoordinate, endCoordinate := sudoku.GetCoordinatesOfHouse(4)

	assert.Equal(t, sudoku.Coordinate{X: 3, Y: 3}, startCoordinate)
	assert.Equal(t, sudoku.Coordinate{X: 5, Y: 5}, endCoordinate)
}

func Test_fills_cell_with_the_number(t *testing.T) {
	board := sudoku.Board{
		{6, 0, 0, 0, 2, 3, 0, 0, 0},
		{7, 4, 0, 1, 0, 0, 0, 0, 9},
		{0, 0, 1, 4, 0, 0, 6, 2, 8},
		{0, 9, 5, 6, 0, 0, 0, 0, 0},
		{0, 1, 0, 2, 3, 0, 0, 6, 0},
		{2, 0, 0, 5, 0, 0, 8, 4, 0},
		{1, 0, 6, 0, 0, 0, 4, 5, 0},
		{0, 0, 0, 0, 4, 1, 0, 8, 0},
		{0, 7, 3, 8, 5, 0, 0, 0, 2},
	}
	expectedBoard := sudoku.Board{
		{6, 0, 0, 0, 2, 3, 0, 0, 0},
		{7, 4, 0, 1, 0, 0, 0, 0, 9},
		{0, 0, 1, 4, 0, 0, 6, 2, 8},
		{0, 9, 5, 6, 8, 0, 0, 0, 0},
		{0, 1, 0, 2, 3, 0, 0, 6, 0},
		{2, 0, 0, 5, 0, 0, 8, 4, 0},
		{1, 0, 6, 0, 0, 0, 4, 5, 0},
		{0, 0, 0, 0, 4, 1, 0, 8, 0},
		{0, 7, 3, 8, 5, 0, 0, 0, 2},
	}

	assert.Equal(t, expectedBoard, board.Fill(sudoku.Coordinate{X: 3, Y: 4}, 8))
}
