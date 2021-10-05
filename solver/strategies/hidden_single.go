package strategies

import (
	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
	"github.com/lourenci/sudoku/utils"
)

type HiddenSingle struct{}

func (HiddenSingle) Find(annotations solver.Annotations) []solver.NumberHint {
	var hiddenSingles []solver.NumberHint

	hiddenSingles = append(hiddenSingles, findHiddenSinglesInHouses(annotations)...)
	hiddenSingles = append(hiddenSingles, findHiddenSinglesInRows(annotations)...)
	hiddenSingles = append(hiddenSingles, findHiddenSinglesInCols(annotations)...)

	return dedupHiddenSingles(hiddenSingles)
}

func findHiddenSinglesInHouses(annotations solver.Annotations) []solver.NumberHint {
	var findedNumbers []solver.NumberHint

	for houseNumber := 0; houseNumber <= 8; houseNumber++ {
		houseAnnotations := annotations.GetAnnotationsFromHouse(houseNumber)

		findedNumbers = append(findedNumbers, findHiddenSinglesInAnnotations(houseAnnotations)...)
	}
	return findedNumbers
}

func findHiddenSinglesInRows(annotations solver.Annotations) []solver.NumberHint {
	var findedNumbers []solver.NumberHint

	for rowNumber := 0; rowNumber <= 8; rowNumber++ {
		rowAnnotations := make(solver.Annotations)
		rowAnnotations[rowNumber] = annotations[rowNumber]

		findedNumbers = append(findedNumbers, findHiddenSinglesInAnnotations(rowAnnotations)...)
	}
	return findedNumbers
}

func findHiddenSinglesInCols(annotations solver.Annotations) []solver.NumberHint {
	var findedNumbers []solver.NumberHint

	for colNumber := 0; colNumber <= 8; colNumber++ {
		colAnnotations := annotations.GetAnnotationsFromCol(colNumber)

		findedNumbers = append(findedNumbers, findHiddenSinglesInAnnotations(colAnnotations)...)
	}
	return findedNumbers
}

func dedupHiddenSingles(hiddenSingles []solver.NumberHint) []solver.NumberHint {
	var dedupedFindedNumbers []solver.NumberHint

	for _, hint := range hiddenSingles {
		if !some(dedupedFindedNumbers, hint) {
			dedupedFindedNumbers = append(dedupedFindedNumbers, hint)
		}
	}

	return dedupedFindedNumbers
}

func some(dedupedFindedNumbers []solver.NumberHint, hint solver.NumberHint) bool {
	for _, findedNumber := range dedupedFindedNumbers {
		if hint.Coordinate.X == findedNumber.Coordinate.X && hint.Coordinate.Y == findedNumber.Coordinate.Y {
			return true
		}
	}
	return false
}

func findHiddenSinglesInAnnotations(annotations solver.Annotations) []solver.NumberHint {
	var findedNumbers []solver.NumberHint

	for i, row := range annotations {
		for j, cell := range row {
			for _, annotationsInCell := range cell {
				if !isNumberAnnotatedInOtherCells(annotationsInCell, annotations) {
					findedNumbers = append(findedNumbers,
						solver.NumberHint{
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
