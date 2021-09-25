package board_test

import (
	"testing"

	board "github.com/lourenci/sudoku/src"
	"github.com/stretchr/testify/assert"
)


func Test_get_house_number_of_a_cell(t *testing.T) {
	assert.Equal(t, 0, board.GetHouseNumberOfCell(0, 0))
	assert.Equal(t, 0, board.GetHouseNumberOfCell(2, 2))
	assert.Equal(t, 4, board.GetHouseNumberOfCell(3, 3))
	assert.Equal(t, 4, board.GetHouseNumberOfCell(5, 5))
	assert.Equal(t, 8, board.GetHouseNumberOfCell(6, 6))
	assert.Equal(t, 8, board.GetHouseNumberOfCell(8, 8))
}
