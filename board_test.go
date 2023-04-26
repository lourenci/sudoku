package sudoku_test

import (
	"testing"

	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/modules/assert"
)

func TestParsedBoard(t *testing.T) {
	t.Run("it parses the string to a board", func(t *testing.T) {
		assert.Equals(
			t,
			sudoku.ParsedBoard(`
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
			sudoku.ParsedBoard("12345678923"),
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
			sudoku.ParsedBoard(`
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
			sudoku.ParsedBoard(`
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
			sudoku.ParsedBoard(`
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

func TestHouse(t *testing.T) {
	t.Run("it returns the board's house", func(t *testing.T) {
		board := sudoku.ParsedBoard(`
			123 456 789
			123 406 789
			123 456 789

			123 456 789
			123 456 789
			123 456 789

			123 456 789
			123 456 789
			987 654 301
		`)
		assert.Equals(t, board.House(sudoku.NewCoordinate(6, 6)), []int{7, 8, 9, 7, 8, 9, 3, 1})
		assert.Equals(t, board.House(sudoku.NewCoordinate(8, 8)), []int{7, 8, 9, 7, 8, 9, 3, 1})
		assert.Equals(t, board.House(sudoku.NewCoordinate(5, 5)), []int{4, 5, 6, 4, 5, 6, 4, 5, 6})
		assert.Equals(t, board.House(sudoku.NewCoordinate(3, 3)), []int{4, 5, 6, 4, 5, 6, 4, 5, 6})
		assert.Equals(t, board.House(sudoku.NewCoordinate(0, 0)), []int{1, 2, 3, 1, 2, 3, 1, 2, 3})
		assert.Equals(t, board.House(sudoku.NewCoordinate(2, 2)), []int{1, 2, 3, 1, 2, 3, 1, 2, 3})
		assert.Equals(t, board.House(sudoku.NewCoordinate(8, 1)), []int{1, 2, 3, 1, 2, 3, 9, 8, 7})
	})
}

func TestAnnotate(t *testing.T) {
	t.Run("it returns the possible numbers for each of the non-filled cell", func(t *testing.T) {
		board := sudoku.ParsedBoard(`
			040 908 010
			000 720 000
			090 006 007

			004 095 070
			009 080 350
			600 317 294

			100 870 060
			000 143 020
			000 560 080
		`)
		assert.Equals(t, board.Annotations(), map[sudoku.Coordinate][]int{
			sudoku.NewCoordinate(0, 0): {2, 3, 5, 7},
			sudoku.NewCoordinate(0, 2): {2, 3, 5, 6, 7},
			sudoku.NewCoordinate(0, 4): {3, 5},
			sudoku.NewCoordinate(0, 6): {5, 6},
			sudoku.NewCoordinate(0, 8): {2, 3, 5, 6},

			sudoku.NewCoordinate(1, 0): {3, 5, 8},
			sudoku.NewCoordinate(1, 1): {1, 3, 5, 6, 8},
			sudoku.NewCoordinate(1, 2): {1, 3, 5, 6, 8},
			sudoku.NewCoordinate(1, 5): {1, 4},
			sudoku.NewCoordinate(1, 6): {4, 5, 6, 8, 9},
			sudoku.NewCoordinate(1, 7): {3, 4},
			sudoku.NewCoordinate(1, 8): {3, 5, 6, 8, 9},

			sudoku.NewCoordinate(2, 0): {2, 3, 5, 8},
			sudoku.NewCoordinate(2, 2): {1, 2, 3, 5, 8},
			sudoku.NewCoordinate(2, 3): {4},
			sudoku.NewCoordinate(2, 4): {3, 5},
			sudoku.NewCoordinate(2, 6): {4, 5, 8},
			sudoku.NewCoordinate(2, 7): {3, 4},

			sudoku.NewCoordinate(3, 0): {2, 3, 8},
			sudoku.NewCoordinate(3, 1): {1, 2, 3, 8},
			sudoku.NewCoordinate(3, 3): {2, 6},
			sudoku.NewCoordinate(3, 6): {1, 6, 8},
			sudoku.NewCoordinate(3, 8): {1, 6, 8},

			sudoku.NewCoordinate(4, 0): {2, 7},
			sudoku.NewCoordinate(4, 1): {1, 2, 7},
			sudoku.NewCoordinate(4, 3): {2, 4, 6},
			sudoku.NewCoordinate(4, 5): {2, 4},
			sudoku.NewCoordinate(4, 8): {1, 6},

			sudoku.NewCoordinate(5, 1): {5, 8},
			sudoku.NewCoordinate(5, 2): {5, 8},

			sudoku.NewCoordinate(6, 1): {2, 3, 5},
			sudoku.NewCoordinate(6, 2): {2, 3, 5},
			sudoku.NewCoordinate(6, 5): {2, 9},
			sudoku.NewCoordinate(6, 6): {4, 5, 9},
			sudoku.NewCoordinate(6, 8): {3, 5, 9},

			sudoku.NewCoordinate(7, 0): {5, 7, 8, 9},
			sudoku.NewCoordinate(7, 1): {5, 6, 7, 8},
			sudoku.NewCoordinate(7, 2): {5, 6, 7, 8},
			sudoku.NewCoordinate(7, 6): {5, 7, 9},
			sudoku.NewCoordinate(7, 8): {5, 9},

			sudoku.NewCoordinate(8, 0): {2, 3, 4, 7, 9},
			sudoku.NewCoordinate(8, 1): {2, 3, 7},
			sudoku.NewCoordinate(8, 2): {2, 3, 7},
			sudoku.NewCoordinate(8, 5): {2, 9},
			sudoku.NewCoordinate(8, 6): {1, 4, 7, 9},
			sudoku.NewCoordinate(8, 8): {1, 3, 9},
		})
	})
}

func TestFill(t *testing.T) {
	t.Run("fills a board's cell with a number", func(t *testing.T) {
		board := sudoku.ParsedBoard(`
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
		board.Fill(sudoku.NewCoordinate(5, 5), 9)

		expectedBoard := sudoku.ParsedBoard(`
			123 456 789
			123 456 789
			123 456 789

			123 456 789
			123 456 789
			123 459 789

			123 456 789
			123 456 789
			123 456 789
		`)

		assert.Equals(
			t,
			board,
			expectedBoard,
		)
	})

	t.Run("panics when an invalid board number is given", func(t *testing.T) {
		assert.NotPanics(t, func() {
			board := sudoku.ParsedBoard(`
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
			board.Fill(sudoku.NewCoordinate(5, 5), 9)
		})
		assert.PanicsWithMessage(t, func() {
			board := sudoku.ParsedBoard(`
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
			board.Fill(sudoku.NewCoordinate(5, 5), 10)
		}, "invalid number")

		assert.NotPanics(t, func() {
			board := sudoku.ParsedBoard(`
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
			board.Fill(sudoku.NewCoordinate(5, 5), 1)
		})
		assert.PanicsWithMessage(t, func() {
			board := sudoku.ParsedBoard(`
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
			board.Fill(sudoku.NewCoordinate(5, 5), 0)
		}, "invalid number")
	})
}
