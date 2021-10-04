package strategies

import (
	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
	"github.com/lourenci/sudoku/utils"
)

type HiddenSingle struct{}

func (h HiddenSingle) Find(annotations solver.Annotations) []solver.Hint {
	var hiddenSingles []solver.Hint

	hiddenSingles = append(hiddenSingles, findHiddenSinglesInHouses(annotations)...)
	hiddenSingles = append(hiddenSingles, findHiddenSinglesInRows(annotations)...)
	hiddenSingles = append(hiddenSingles, findHiddenSinglesInCols(annotations)...)

	return dedupHiddenSingles(hiddenSingles)
}

func findHiddenSinglesInHouses(annotations solver.Annotations) []solver.Hint {
	var findedNumbers []solver.Hint

	for houseNumber := 0; houseNumber <= 8; houseNumber++ {
		houseAnnotations := annotations.GetAnnotationsFromHouse(houseNumber)

		findedNumbers = append(findedNumbers, findHiddenSinglesInAnnotations(houseAnnotations)...)
	}
	return findedNumbers
}

func findHiddenSinglesInRows(annotations solver.Annotations) []solver.Hint {
	var findedNumbers []solver.Hint

	for rowNumber := 0; rowNumber <= 8; rowNumber++ {
		rowAnnotations := make(solver.Annotations)
		rowAnnotations[rowNumber] = annotations[rowNumber]

		findedNumbers = append(findedNumbers, findHiddenSinglesInAnnotations(rowAnnotations)...)
	}
	return findedNumbers
}

func findHiddenSinglesInCols(annotations solver.Annotations) []solver.Hint {
	var findedNumbers []solver.Hint

	for colNumber := 0; colNumber <= 8; colNumber++ {
		colAnnotations := annotations.GetAnnotationsFromCol(colNumber)

		findedNumbers = append(findedNumbers, findHiddenSinglesInAnnotations(colAnnotations)...)
	}
	return findedNumbers
}

func dedupHiddenSingles(hiddenSingles []solver.Hint) []solver.Hint {
	var dedupedFindedNumbers []solver.Hint

	for _, hint := range hiddenSingles {
		if !some(dedupedFindedNumbers, hint) {
			dedupedFindedNumbers = append(dedupedFindedNumbers, hint)
		}
	}

	return dedupedFindedNumbers
}

func some(dedupedFindedNumbers []solver.Hint, hint solver.Hint) bool {
	for _, findedNumber := range dedupedFindedNumbers {
		if hint.X == findedNumber.X && hint.Y == findedNumber.Y {
			return true
		}
	}
	return false
}

func findHiddenSinglesInAnnotations(annotations solver.Annotations) []solver.Hint {
	var findedNumbers []solver.Hint

	for i, row := range annotations {
		for j, cell := range row {
			for _, annotationsInCell := range cell {
				if !isNumberAnnotatedInOtherCells(annotationsInCell, annotations) {
					findedNumbers = append(findedNumbers,
						solver.Hint{
							Coordinate: sudoku.Coordinate{X: i, Y: j},
							Number:     annotationsInCell,
						},
					)
				}
			}
		}
	}

	return findedNumbers
}

func isNumberAnnotatedInOtherCells(number int, annotations solver.Annotations) bool {
	qtyAppearance := 0

	for _, row := range annotations {
		for _, cell := range row {
			if utils.FindIndex(cell, number) != nil {
				qtyAppearance++
				if qtyAppearance > 1 {
					return true
				}
			}
		}
	}

	return false
}
