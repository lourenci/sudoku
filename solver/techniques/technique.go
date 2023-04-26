package techniques

import "github.com/lourenci/sudoku"

type Board interface {
	Annotations() map[sudoku.Coordinate][]int
	Fill(sudoku.Coordinate, int)
}

type Technique interface {
	Name() string
	Solve(Board)
}
