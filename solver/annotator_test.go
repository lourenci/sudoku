package solver_test

import (
	"testing"

	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
	"github.com/stretchr/testify/assert"
)

func Test_annotates_with_the_candidates(t *testing.T) {
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

	expectedAnnotations := solver.Annotations{
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

	assert.Equal(t, expectedAnnotations, solver.NewAnnotation().Annotate(board))
}

func Test_get_annotations_in_a_given_house(t *testing.T) {
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
			6: []int{5, 8},
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

	expectedAnnotations := solver.Annotations{
		3: {
			3: []int{4, 6, 7},
			4: []int{4, 6},
		},
		5: {
			3: []int{2, 4, 6},
			4: []int{4, 6},
			5: []int{6, 8},
		},
	}

	assert.Equal(t, expectedAnnotations, annotations.GetAnnotationsFromHouse(4))
}

func Test_get_annotations_in_a_given_col(t *testing.T) {
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
			6: []int{5, 8},
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

	expectedAnnotations := solver.Annotations{
		2: {
			4: []int{3, 4, 6},
		},
		3: {
			4: []int{4, 6},
		},
		5: {
			4: []int{4, 6},
		},
		6: {
			4: []int{1, 3, 6},
		},
	}

	assert.Equal(t, expectedAnnotations, annotations.GetAnnotationsFromCol(4))
}

func Test_fills_the_coordinates_with_the_given_notes(t *testing.T) {
	annotations := solver.Annotations{
		8: {
			0: []int{4, 9},
			7: []int{1, 9},
		},
	}

	annotations.Fill(sudoku.Coordinate{X: 8, Y: 1}, []int{4, 9})
	assert.Equal(t, solver.Annotations{
		8: {
			0: []int{4, 9},
			1: []int{4, 9},
			7: []int{1, 9},
		},
	}, annotations)

	annotations.Fill(sudoku.Coordinate{X: 8, Y: 7}, []int{4, 9})
	assert.Equal(t, solver.Annotations{
		8: {
			0: []int{4, 9},
			1: []int{4, 9},
			7: []int{4, 9},
		},
	}, annotations)

	annotations.Fill(sudoku.Coordinate{X: 0, Y: 0}, []int{4, 9})
	assert.Equal(t, solver.Annotations{
		0: {
			0: []int{4,9},
		},
		8: {
			0: []int{4, 9},
			1: []int{4, 9},
			7: []int{4, 9},
		},
	}, annotations)
}
