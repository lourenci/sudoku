package strategies

import (
	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
	"github.com/lourenci/sudoku/utils"
)

type HiddenSingle struct{}

func (h HiddenSingle) Find(annotations solver.Annotations) []solver.Hint {
	var findedNumbers []solver.Hint

	for houseNumber := 0; houseNumber <= 8; houseNumber++ {
		houseAnnotations := annotations.GetAnnotationsFromHouse(houseNumber)

		for i, row := range houseAnnotations {
			for j, cell := range row {
				for _, number := range cell {
					if !isNumberInOtherCellsOfHouse(number, houseAnnotations) {
						findedNumbers = append(findedNumbers, solver.Hint{Coordinate: sudoku.Coordinate{X: i, Y: j}, Number: number})
					}
				}
			}
		}
	}

	for colNumber := 0; colNumber <= 8; colNumber++ {
		colAnnotations := annotations.GetAnnotationsFromCol(colNumber)

		for i, row := range colAnnotations {
			for j, cell := range row {
				for _, number := range cell {
					if !isNumberInOtherCellsOfCol(number, colAnnotations) {
						findedNumbers = append(findedNumbers, solver.Hint{Coordinate: sudoku.Coordinate{X: i, Y: j}, Number: number})
					}
				}
			}
		}
	}

	for rowNumber := 0; rowNumber <= 8; rowNumber++ {
		rowAnnotations := make(solver.Annotations)
		rowAnnotations[rowNumber] = annotations[rowNumber]

		for i, row := range rowAnnotations {
			for j, cell := range row {
				for _, number := range cell {
					if !isNumberInOtherCellsOfRow(number, rowAnnotations) {
						findedNumbers = append(findedNumbers, solver.Hint{Coordinate: sudoku.Coordinate{X: i, Y: j}, Number: number})
					}
				}
			}
		}
	}

	var dedupedFindedNumbers []solver.Hint

	for _, hint := range findedNumbers {
		if !newFunction(dedupedFindedNumbers, hint) {
			dedupedFindedNumbers = append(dedupedFindedNumbers, hint)
		}
	}

	return dedupedFindedNumbers
}

func newFunction(dedupedFindedNumbers []solver.Hint, hint solver.Hint) bool {
	for _, findedNumber := range dedupedFindedNumbers {
		if hint.X == findedNumber.X && hint.Y == findedNumber.Y {
			return true
		}
	}
	return false
}

func isNumberInOtherCellsOfHouse(number int, houseAnnotation solver.Annotations) bool {
	qtyAppearance := 0

	for _, row := range houseAnnotation {
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

func isNumberInOtherCellsOfCol(number int, colAnnotation solver.Annotations) bool {
	qtyAppearance := 0

	for _, row := range colAnnotation {
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

func isNumberInOtherCellsOfRow(number int, rowAnnotation solver.Annotations) bool {
	qtyAppearance := 0

	for _, row := range rowAnnotation {
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
