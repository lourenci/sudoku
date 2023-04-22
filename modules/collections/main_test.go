package collections_test

import (
	"testing"

	"lourenci.com/sudoku/modules/assert"
	"lourenci.com/sudoku/modules/collections"
)

func TestExcept(t *testing.T) {
	assert.Equals(t, collections.Except([]int{1, 2, 3}, []int{}), []int{1, 2, 3})
	assert.Equals(t, collections.Except([]int{1, 2, 3}, []int{4}), []int{1, 2, 3})
	assert.Equals(t, collections.Except([]int{1, 2, 3}, []int{3}), []int{1, 2})
}