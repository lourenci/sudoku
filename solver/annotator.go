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

func (a Annotations) GetAnnotationsFromCol(colNumber int) Annotations {
	annotations := make(Annotations)

	for i, row := range a {
		if row[colNumber] != nil {
			annotations[i] = make(map[int][]int)
			annotations[i][colNumber] = row[colNumber]
		}
	}

	return annotations
}

func (a Annotations) Fill(coordinate sudoku.Coordinate, numbers []int) {
	if a[coordinate.X] == nil {
		a[coordinate.X] = make(map[int][]int)
	}

	a[coordinate.X][coordinate.Y] = numbers
}

func (a Annotations) annotateCellWithMissingNumbers(b sudoku.Board, c sudoku.Coordinate) {
	missingNumbers := b.MissingNumbersInCell(c)
	if missingNumbers != nil {
		a[c.X][c.Y] = missingNumbers
	}
}

func (a Annotations) GetAnnotationsFromHouse(houseNumber int) Annotations {
	startCoordinate, endCoordinate := sudoku.GetCoordinatesOfHouse(houseNumber)

	return a.getAnnotationsInRange(startCoordinate, endCoordinate)
}

func (a Annotations) getAnnotationsInRange(startCoordinate sudoku.Coordinate, endCoordinate sudoku.Coordinate) Annotations {
	annotations := make(Annotations)
	for i := startCoordinate.X; i <= endCoordinate.X; i++ {
		for j := startCoordinate.Y; j <= endCoordinate.Y; j++ {
			if a[i][j] != nil {
				if annotations[i] == nil {
					annotations[i] = make(map[int][]int)
				}

				annotations[i][j] = a[i][j]
			}
		}
	}
	return annotations
}
