package board_test

import (
	"testing"

	board "github.com/lourenci/sudoku/src"
	src "github.com/lourenci/sudoku/src"
	"github.com/stretchr/testify/assert"
)

func Test_annotates_with_the_candidates(t *testing.T) {
	board := board.Board{
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

	expectedAnnotations := src.Annotations{
		0: {
			1: []int{5, 8},
			2: []int{8, 9},
			3: []int{7, 9},
			6: []int{1, 5, 7},
			7: []int{1, 7},
			8: []int{1, 4, 5, 7},
		},
		1: {
			2: []int{2, 8},
			4: []int{6, 8},
			5: []int{5, 6, 8},
			6: []int{3, 5},
			7: []int{3},
		},
		2: {
			0: []int{3, 5, 9},
			1: []int{3, 5},
			4: []int{7, 9},
			5: []int{5, 7, 9},
		},
		3: {
			0: []int{3, 4, 8},
			4: []int{1, 7, 8},
			5: []int{4, 7, 8},
			6: []int{1, 2, 3, 7},
			7: []int{1, 3, 7},
			8: []int{1, 3, 7},
		},
		4: {
			0: []int{4, 8},
			2: []int{4, 7, 8},
			5: []int{4, 7, 8, 9},
			6: []int{5, 7, 9},
			8: []int{5, 7},
		},
		5: {
			1: []int{3, 6},
			2: []int{7},
			4: []int{1, 7, 9},
			5: []int{7, 9},
			8: []int{1, 3, 7},
		},
		6: {
			1: []int{2, 8},
			3: []int{3, 7, 9},
			4: []int{7, 9},
			5: []int{2, 7, 9},
			8: []int{3, 7},
		},
		7: {
			0: []int{5, 9},
			1: []int{2, 5},
			2: []int{2, 9},
			3: []int{3, 7, 9},
			6: []int{3, 7, 9},
			8: []int{3, 6, 7},
		},
		8: {
			0: []int{4, 9},
			5: []int{6, 9},
			6: []int{1, 9},
			7: []int{1, 9},
		},
	}

	assert.Equal(t, expectedAnnotations, board.Annotate())
}

func Test_get_house_number_of_a_cell(t *testing.T) {
	assert.Equal(t, 0, board.GetHouseOfACell(0, 0))
	assert.Equal(t, 0, board.GetHouseOfACell(2, 2))
	assert.Equal(t, 4, board.GetHouseOfACell(3, 3))
	assert.Equal(t, 4, board.GetHouseOfACell(5, 5))
	assert.Equal(t, 8, board.GetHouseOfACell(6, 6))
	assert.Equal(t, 8, board.GetHouseOfACell(8, 8))
}
