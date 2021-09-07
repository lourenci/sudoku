package board_test

import (
	"testing"

	src "github.com/lourenci/sudoku/src"
	"github.com/stretchr/testify/assert"
)

func Test_annotates_the_lines_with_the_candidates(t *testing.T) {
	board := src.Board{
		[9]int{6, 0, 0, 0, 2, 3, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	expectedAnnotations := src.Annotations{0: {1: []int{5, 8}, 2: []int{8, 9}, 3: []int{7, 9}, 6: []int{1, 5, 7}, 7: []int{1, 7}, 8: []int{1, 4, 5, 7}}}

	assert.Equal(t, expectedAnnotations, board.Annotate())
}

func Test_annotates_with_the_candidates(t *testing.T) {
	board := src.Board{
		[9]int{6, 0, 0, 0, 2, 3, 0, 0, 0},
		[9]int{7, 4, 0, 1, 0, 0, 0, 0, 9},
		[9]int{0, 0, 1, 4, 0, 0, 6, 2, 8},
		[9]int{0, 9, 5, 6, 0, 0, 0, 0, 0},
		[9]int{0, 1, 0, 2, 3, 0, 0, 6, 0},
		[9]int{2, 0, 0, 5, 0, 8, 4, 0, 0},
		[9]int{1, 0, 6, 0, 0, 0, 4, 5, 0},
		[9]int{0, 0, 0, 0, 4, 1, 0, 8, 0},
		[9]int{0, 7, 3, 8, 5, 0, 0, 0, 2},
	}

	expectedAnnotations := src.Annotations{0: {1: []int{5, 8}, 2: []int{8, 9}, 3: []int{7, 9}, 6: []int{1, 5, 7}, 7: []int{1, 7}, 8: []int{1, 4, 5, 7}}}

	assert.Equal(t, expectedAnnotations, board.Annotate())
}
