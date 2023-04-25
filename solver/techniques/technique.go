package techniques

import "github.com/lourenci/sudoku"

type Board interface {
	Numbers() [9][9]int
	Annotations() map[sudoku.Coordinate][]int
	Fill(sudoku.Coordinate, int)
}

type Technique interface {
	Name() string
	Solve(Board)
}
