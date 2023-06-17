package techniques

import "github.com/lourenci/sudoku"

type HiddenSingle struct{}

func (HiddenSingle) Name() string {
	return "Hidden Single"
}

func (HiddenSingle) Solve(board Board) {
	for c, annotations := range board.Annotations() {
		for _, annotation := range annotations {
			howManyOfAnnotation := 0

			for j := 0; j < 9 && howManyOfAnnotation <= 1; j++ {
				x := board.Annotations()[sudoku.NewCoordinate(c.X, j)]
				for _, v2 := range x {
					if annotation == v2 {
						howManyOfAnnotation++
						break
					}
				}
			}

			if howManyOfAnnotation == 1 {
				board.Fill(sudoku.NewCoordinate(c.X, c.Y), annotation)
			}
		}
	}
}
