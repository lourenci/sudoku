package solver

import "github.com/lourenci/sudoku"

type Strategy interface {
	Find(Annotations) []NumberHint
}

type NumberHint struct {
	Coordinate sudoku.Coordinate
	Number int
}

type BoardSector int

const (
	Row BoardSector = iota
	Col
	House
)

type AnnotationHint struct {
	Coordinate sudoku.Coordinate
	BoardSector
}

type Solve struct {
	strategies []Strategy
}

func NewSolve(strategies []Strategy) Solve {
	return Solve{strategies}
}

func (s Solve) Solve(b sudoku.Board) sudoku.Board {
	return s.solve(b, sudoku.Board{})
}

func (s Solve) solve(b sudoku.Board, previousBoard sudoku.Board) sudoku.Board {
	if b == previousBoard {
		return b
	}

	previousBoard = b

	annotations := NewAnnotation().Annotate(b)

	hints := s.strategies[0].Find(annotations)

	for _, hint := range hints {
		b[hint.Coordinate.X][hint.Coordinate.Y] = hint.Number
	}

	return s.solve(b, previousBoard)
}
