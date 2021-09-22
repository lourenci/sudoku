package board_test

import (
	"testing"

	board "github.com/lourenci/sudoku/src"
	"github.com/stretchr/testify/assert"
)

func Test_finds_and_returns_the_index_of_an_element(t *testing.T) {
	numbers := [9]int{1, 3, 5, 7, 9}

	assert.Equal(t, 2, *board.FindIndex(numbers[:], 5))
	assert.Nil(t, board.FindIndex(numbers[:], 8))
}
