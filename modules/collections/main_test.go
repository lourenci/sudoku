package collections_test

import (
	"testing"

	"github.com/lourenci/sudoku/modules/assert"
	"github.com/lourenci/sudoku/modules/collections"
)

func TestExcept(t *testing.T) {
	assert.Equals(t, collections.Except([]int{1, 2, 3}, []int{}), []int{1, 2, 3})
	assert.Equals(t, collections.Except([]int{1, 2, 3}, []int{4}), []int{1, 2, 3})
	assert.Equals(t, collections.Except([]int{1, 2, 3}, []int{3}), []int{1, 2})
}

func TestIntersect(t *testing.T) {
	assert.Equals(t, collections.Intersect([]int{1, 2, 3}, []int{}), []int{})
	assert.Equals(t, collections.Intersect([]int{1, 2, 3}, []int{4}), []int{})
	assert.Equals(t, collections.Intersect([]int{1, 2, 3}, []int{2, 3}), []int{2, 3})
}
