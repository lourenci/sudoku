package strategies

import (
	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
)

type NakedPair struct{}

func (h NakedPair) Find(annotations solver.Annotations) []solver.AnnotationHint {
	var foundNakedPairs []solver.AnnotationHint

	for x, rowNumber := range annotations {
		possibleNakedPair := make(map[int][]int)

		for y, cell := range rowNumber {
			if len(cell) == 2 {
				for naked_c, naked := range possibleNakedPair {
					if naked[0] == cell[0] && naked[1] == cell[1] {
						a := solver.NewAnnotation()
						a.Fill(sudoku.Coordinate{X: x, Y: naked_c}, cell)
						a.Fill(sudoku.Coordinate{X: x, Y: y}, cell)
						foundNakedPairs = append(foundNakedPairs, solver.AnnotationHint{Annotations: a})
					}
				}

				possibleNakedPair[y] = append(possibleNakedPair[y], cell...)
			}
		}

	}

	return foundNakedPairs
}
