package utils_test

import (
	"testing"

	"github.com/lourenci/sudoku/utils"
	"github.com/stretchr/testify/assert"
)

func Test_finds_and_returns_the_index_of_an_element(t *testing.T) {
	numbers := [9]int{1, 3, 5, 7, 9}

	assert.Equal(t, 2, *utils.FindIndex(numbers[:], 5))
	assert.Nil(t, utils.FindIndex(numbers[:], 8))
}
