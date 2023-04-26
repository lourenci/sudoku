package techniques_test

import (
	"testing"

	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/modules/assert"
	"github.com/lourenci/sudoku/solver/techniques"
)

func TestNakedSingle(t *testing.T) {
	t.Run("it solves the board using the naked single technique", func(t *testing.T) {
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
		techniques.NakedSingle{}.Solve(&board)

		expectedBoard := sudoku.ParsedBoard(`
			040 908 010
			000 720 000
			090 406 007

			004 095 070
			009 080 350
			600 317 294

			100 870 060
			000 143 020
			000 560 080
		`)

		assert.Equals(
			t,
			board,
			expectedBoard,
		)
	})
}
