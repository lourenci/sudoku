package solver

import sudoku "github.com/lourenci/sudoku/src"

func Solve(b sudoku.Board) Board {
	annotations := NewAnnotation().Annotate(b)

	findeds := NewStrategy(annotations).Find()

	for _, coordinate := range findeds {
		b[coordinate.X][coordinate.Y] = coordinate.Number
	}

	if b.IsComplete() {
		return b
	}

	return Solve(b)
}
