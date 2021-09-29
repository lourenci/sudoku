package solver_test

import (
	"testing"

	"github.com/lourenci/sudoku/src/solver"
	"github.com/stretchr/testify/assert"
)

func Test_finds_for_only_one_missing_number(t *testing.T) {
	annotations := solver.Annotations{
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

	assert.ElementsMatch(t, []solver.Coordinate{{1, 7, 3}, {5, 2, 7}}, solver.HiddenSingle{}.Find(annotations))
}
