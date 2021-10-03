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
		houseAnnotation := annotations.GetAnnotationsFromHouse(houseNumber)

		for i, row := range houseAnnotation {
			for j, cell := range row {
				for _, number := range cell {
					if !isNumberInOtherCellsOfHouse(number, houseAnnotation) {
						findedNumbers = append(findedNumbers, solver.Hint{Coordinate: sudoku.Coordinate{X: i, Y: j}, Number: number})
					}
				}
			}
		}
	}

	return findedNumbers
}

func isNumberInOtherCellsOfHouse(number int, houseAnnotation solver.Annotations) bool {
	qtyAppareance := 0

	for _, row := range houseAnnotation {
		for _, cell := range row {
			if utils.FindIndex(cell, number) != nil {
				qtyAppareance++
				if qtyAppareance > 1 {
					return true
				}
			}
		}
	}

	return false
}
