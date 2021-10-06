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

	hints := s.findHintsUsingTheAnnotations(annotations)

	return s.tryToSolve(fillBoardUsingTheHints(hints, b), previousBoard)
}

func fillBoardUsingTheHints(hints []Hint, b sudoku.Board) sudoku.Board {
	for _, coordinate := range hints {
		b = b.Fill(coordinate.Coordinate, coordinate.Number)
	}

	return b
}

func (s Solve) findHintsUsingTheAnnotations(annotations Annotations) []Hint {
	var hints []Hint

	for _, strategy := range s.strategies {
		hints = append(hints, strategy.Find(annotations)...)
	}
	return hints
}
