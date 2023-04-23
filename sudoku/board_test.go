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
			`),
			sudoku.NewBoard(
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
			sudoku.NewBoard(
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
			sudoku.NewBoard(
				[9][9]int{
					{1, 2, 3, 4, 5, 6, 7, 8, 9},
					{2, 3, 5, 0, 0, 2, 3, 3, 0},
				},
			),
		)
	})

	t.Run("it doesn't accept more than 81 numbers (a 9x9 board)", func(t *testing.T) {
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
			sudoku.NewBoard(
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

func TestColumn(t *testing.T) {
	t.Run("it returns the board's column", func(t *testing.T) {
		board := sudoku.Parse(`
			123 456 789
			123 456 789
			123 456 789

			123 456 789
			123 456 789
			123 456 789

			123 456 789
			123 456 789
			123 456 789
		`)
		assert.Equals(t, board.Column(0), [9]int{1, 1, 1, 1, 1, 1, 1, 1, 1})
		assert.Equals(t, board.Column(8), [9]int{9, 9, 9, 9, 9, 9, 9, 9, 9})
	})
}

func TestRow(t *testing.T) {
	t.Run("it returns the board's row", func(t *testing.T) {
		board := sudoku.Parse(`
			123 456 789
			123 406 789
			123 456 789

			123 456 789
			123 456 789
			123 456 789

			123 456 789
			123 456 789
			987654321
		`)
		assert.Equals(t, board.Row(0), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		assert.Equals(t, board.Row(1), []int{1, 2, 3, 4, 6, 7, 8, 9})
		assert.Equals(t, board.Row(8), []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
	})
}
