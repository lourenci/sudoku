package techniques_test

import (
	"testing"

	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/modules/assert"
	"github.com/lourenci/sudoku/solver/techniques"
)

func TestHiddenSingle(t *testing.T) {
	t.Run("it solves the board using the hidden single technique (row)", func(t *testing.T) {
		board := sudoku.ParsedBoard(`
			100 030 000
			090 600 700
			006 000 001

			630 000 805
			080 050 040
			700 000 010

			000 000 100
			007 005 020
			000 920 034
		`)
		techniques.HiddenSingle{}.Solve(&board)

		expectedBoard := sudoku.ParsedBoard(`
			100 030 000
			090 600 700
			006 000 001

			630 000 805
			080 050 040
			700 000 010

			000 000 100
			007 005 020
			000 927 034
		`)

		assert.Equals(
			t,
			board,
			expectedBoard,
		)

		board = sudoku.ParsedBoard(`
			500 400 781
			008 000 900
			030 000 002

			000 800 006
			200 030 009
			003 701 000

			100 000 090
			005 000 200
			067 004 008
		`)
		techniques.HiddenSingle{}.Solve(&board)

		expectedBoard = sudoku.ParsedBoard(`
			500 403 781
			008 000 900
			030 000 002

			000 800 006
			200 030 009
			003 701 000

			100 000 090
			005 000 200
			067 004 008
		`)

		assert.Equals(
			t,
			board,
			expectedBoard,
		)
	})
}
