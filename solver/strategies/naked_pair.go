package strategies

import (
	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
)

type NakedPair struct{}

func (h NakedPair) Find(annotations solver.Annotations) []solver.AnnotationHint {
	var foundNakedPairs []solver.AnnotationHint

	foundNakedPairs = append(foundNakedPairs, findNakedPairsInRows(annotations)...)
	foundNakedPairs = append(foundNakedPairs, findNakedPairsInCols(annotations)...)

	return foundNakedPairs
}

func findNakedPairsInRows(annotations solver.Annotations) []solver.AnnotationHint {
	var nakedPairs []solver.AnnotationHint

	for x, row := range annotations {
		possibleNakedPairInRow := make(map[int][]int)

		for y, cell := range row {
			isNakedPair := len(cell) == 2
			if isNakedPair {
				for possibleNakedY, possibleNaked := range possibleNakedPairInRow {
					if possibleNaked[0] == cell[0] && possibleNaked[1] == cell[1] {
						a := solver.NewAnnotation()
						a.Fill(sudoku.Coordinate{X: x, Y: possibleNakedY}, cell)
						a.Fill(sudoku.Coordinate{X: x, Y: y}, cell)
						nakedPairs = append(nakedPairs, solver.AnnotationHint{Annotations: a})
					}
				}

				possibleNakedPairInRow[y] = append(possibleNakedPairInRow[y], cell...)
			}
		}
	}

	return nakedPairs
}

func findNakedPairsInCols(annotations solver.Annotations) []solver.AnnotationHint {
	var nakedPairs []solver.AnnotationHint

	for colNumber := 0; colNumber <= 8; colNumber++ {
		colAnnotations := annotations.GetAnnotationsFromCol(colNumber)

		nakedPairs = append(nakedPairs, findNakedPairsInAnnotations(colAnnotations)...)
	}

	return nakedPairs
}

func findNakedPairsInAnnotations(annotations solver.Annotations) []solver.AnnotationHint {
	var nakedPairs []solver.AnnotationHint
	possibleNakedPairInRow := make(solver.Annotations)

	for x, row := range annotations {
		for y, cell := range row {
			isNakedPair := len(cell) == 2
			if isNakedPair {
				for possibleNakedX, possibleNaked := range possibleNakedPairInRow {
					for possibleNakedY, possibleNaked := range possibleNaked {
						if possibleNaked[0] == cell[0] && possibleNaked[1] == cell[1] {
							a := solver.NewAnnotation()
							a.Fill(sudoku.Coordinate{X: possibleNakedX, Y: possibleNakedY}, cell)
							a.Fill(sudoku.Coordinate{X: x, Y: y}, cell)
							nakedPairs = append(nakedPairs, solver.AnnotationHint{Annotations: a})
						}
					}
				}

				if possibleNakedPairInRow[x] == nil {
					possibleNakedPairInRow[x] = make(map[int][]int)
				}

				possibleNakedPairInRow[x][y] = cell
			}
		}
	}

	return nakedPairs
}
