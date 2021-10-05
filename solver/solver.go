package solver

import "github.com/lourenci/sudoku"

type Strategy interface {
	Find(Annotations) []Hint
}

type Hint struct {
	sudoku.Coordinate
	Number int
}

type Solve struct {
	strategies []Strategy
}

func NewSolve(strategies []Strategy) Solve {
	return Solve{strategies}
}

func (s Solve) Solve(b sudoku.Board) sudoku.Board {
	return s.tryToSolve(b, sudoku.Board{})
}

func (s Solve) tryToSolve(b sudoku.Board, previousBoard sudoku.Board) sudoku.Board {
	if b == previousBoard {
		return b
	}

	previousBoard = b

	annotations := NewAnnotation().Annotate(b)

	var findeds []Hint

	for _, strategy := range s.strategies {
		findeds = append(findeds, strategy.Find(annotations)...)
	}

	for _, coordinate := range findeds {
		b[coordinate.X][coordinate.Y] = coordinate.Number
	}

	return s.tryToSolve(b, previousBoard)
}
