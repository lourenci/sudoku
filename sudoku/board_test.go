package sudoku_test

import (
	"testing"

	"lourenci.com/sudoku/modules/assert"
	"lourenci.com/sudoku/sudoku"
)

func TestParse(t *testing.T) {
	t.Run("it parses the string to a board", func(t *testing.T) {
		assert.Equals(
			t,
			sudoku.Parse(`
				123 456 789
				123 456 789
				123 456 789

				123 456 789
				123 456 789
				123 456 789

				123 456 789
				123 456 789
				123 456 789

				f
			`),
			sudoku.Board(
				[9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
				},
			),
		)
	})

	t.Run("it accepts a non-complete board", func(t *testing.T) {
		assert.Equals(
			t,
			sudoku.Parse("12345678923"),
			sudoku.Board(
				[9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{2, 3, 0, 0, 0, 0, 0, 0, 0},
				},
			),
		)
	})

	t.Run("it ignores non-number", func(t *testing.T) {
		assert.Equals(
			t,
			sudoku.Parse(`
				123 456 789
				235b00 233
			`),
			sudoku.Board(
				[9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{2, 3, 5, 0, 0, 2, 3, 3, 0},
				},
			),
		)
	})

	t.Run("it doesn't accept more than 81 numbers (a 9x9 board)", func(t *testing.T) {
		assert.PanicsWithMessage(t, func() {
			sudoku.Parse(`
				123 456 789
				123 456 789
				123 456 789

				123 456 789
				123 456 789
				123 456 789

				123 456 789
				123 456 789
				123 456 789

				9
			`)
		}, "please, provide a 9x9 board")
	})
}
