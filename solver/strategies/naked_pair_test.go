package strategies_test

import (
	"testing"

	"github.com/lourenci/sudoku/solver"
	"github.com/lourenci/sudoku/solver/strategies"
	"github.com/stretchr/testify/assert"
)

func Test_finds_the_naked_pairs_in_the_same_row(t *testing.T) {
	annotations := solver.Annotations{
		0: {
			0: []int{2, 6, 8},
			3: []int{3, 6},
			5: []int{1, 6},
			7: []int{1, 2, 3, 6, 8},
			8: []int{1, 2, 8},
		},
		1: {
			0: []int{6, 7},
			3: []int{6, 7, 9},
			6: []int{6, 9},
			7: []int{1, 6},
		},
		2: {
			0: []int{2, 6, 7, 8},
			1: []int{2, 7, 8},
			3: []int{3, 4, 6, 7, 9},
			4: []int{3, 4, 6},
			5: []int{6, 7, 9},
			6: []int{6, 8, 9},
			8: []int{2, 8, 9},
		},
		3: {
			0: []int{1, 3, 6, 7, 8, 9},
			1: []int{1, 7, 8},
			3: []int{4, 6, 7},
			4: []int{4, 6},
			6: []int{8, 9},
			7: []int{1, 4, 8},
			8: []int{1, 8, 9},
		},
		4: {
			1: []int{5, 7, 8},
			2: []int{7, 8},
			6: []int{7, 8},
			7: []int{2, 8},
		},
		5: {
			0: []int{1, 6, 8, 9},
			1: []int{1, 5, 8},
			2: []int{6, 8},
			3: []int{2, 4, 6},
			4: []int{4, 6},
			5: []int{6, 8},
			8: []int{1, 2, 5, 8, 9},
		},
		6: {
			0: []int{1, 7, 8},
			1: []int{1, 4, 7, 8},
			2: []int{7, 8},
			3: []int{3, 5, 6, 9},
			4: []int{1, 3, 6},
			5: []int{1, 6, 9},
			7: []int{6, 8},
			8: []int{5, 8},
		},
		8: {
			0: []int{1, 2, 8},
			1: []int{1, 2, 8},
			3: []int{5, 6},
			5: []int{1, 6},
			7: []int{6, 8},
		},
	}

	assert.ElementsMatch(
		t,
		[]solver.AnnotationHint{},
		strategies.NakedPair{}.Find(annotations),
	)
}
