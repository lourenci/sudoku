package solver

import "github.com/lourenci/sudoku"

type Annotations map[int]map[int][]int

func NewAnnotation() Annotations {
	return make(Annotations)
}

func (a Annotations) Annotate(b sudoku.Board) Annotations {
	for x := 0; x <= 8; x++ {
		a[x] = make(map[int][]int)

		for y := 0; y <= 8; y++ {
			a.annotateCellWithMissingNumbers(b, sudoku.Coordinate{X: x, Y: y})
		}
	}
	return a
}

func (a Annotations) annotateCellWithMissingNumbers(b sudoku.Board, c sudoku.Coordinate) {
	missingNumbers := b.MissingNumbersInCell(c)
	if missingNumbers != nil {
		a[c.X][c.Y] = missingNumbers
	}
}
