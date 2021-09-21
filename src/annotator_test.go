package board_test

import (
	"testing"

	board "github.com/lourenci/sudoku/src"
	"github.com/stretchr/testify/assert"
)

func Test_annotates_the_missing_numbers_in_the_line(t *testing.T) {
	expectedAnnotations := map[int][]int{
		1: {1, 4, 5, 7, 8, 9},
		2: {1, 4, 5, 7, 8, 9},
		3: {1, 4, 5, 7, 8, 9},
		6: {1, 4, 5, 7, 8, 9},
		7: {1, 4, 5, 7, 8, 9},
		8: {1, 4, 5, 7, 8, 9},
	}

	l := []int{6, 0, 0, 0, 2, 3, 0, 0, 0}

	assert.Equal(t, expectedAnnotations, board.AnnotateMissingNumbers(l))
}

func Test_returns_the_missing_numbers(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6, 7, 8, 9}, board.FindMissingNumbers([]int{1, 3, 5}))
}
