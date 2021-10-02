package solver

import "github.com/lourenci/sudoku"

type Annotations map[int]map[int][]int

func NewAnnotation() Annotations {
	return make(Annotations)
}

func (a Annotations) Annotate(b sudoku.Board) Annotations {
	for rowNumber := 0; rowNumber <= 8; rowNumber++ {
		a[rowNumber] = make(map[int][]int)

		for colNumber := 0; colNumber <= 8; colNumber++ {
			a.annotateCellWithMissingNumbers(b, rowNumber, colNumber)
		}
	}
	return a
}

func (a Annotations) annotateCellWithMissingNumbers(b sudoku.Board, rowNumber int, colNumber int) {
	missingNumbers := b.MissingNumbersInCell(rowNumber, colNumber)
	if missingNumbers != nil {
		a[rowNumber][colNumber] = missingNumbers
	}
}
