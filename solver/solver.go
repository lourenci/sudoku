package solver

import "github.com/lourenci/sudoku"

type Strategy interface {
	Find(Annotations) []Hint
}

type Hint struct {
	sudoku.Coordinate
	Number int
}

type AnnotationHint struct {
	Annotations
}

type Solver struct {
	strategies []Strategy
}

func NewSolver(strategies []Strategy) Solver {
	return Solver{strategies}
}

func (s Solver) Solve(b sudoku.Board) sudoku.Board {
	return s.tryToSolve(b, sudoku.Board{})
}

func (s Solver) tryToSolve(b sudoku.Board, previousBoard sudoku.Board) sudoku.Board {
	if b == previousBoard {
		return b
	}

	previousBoard = b

	annotations := NewAnnotation().Annotate(b)

	hints := s.findHintsUsingTheAnnotations(annotations)

	return s.tryToSolve(fillBoardUsingTheHints(hints, b), previousBoard)
}

func fillBoardUsingTheHints(hints []Hint, b sudoku.Board) sudoku.Board {
	for _, hint := range hints {
		b = b.Fill(hint.Coordinate, hint.Number)
	}

	return b
}

func (s Solver) findHintsUsingTheAnnotations(annotations Annotations) []Hint {
	var hints []Hint

	for _, strategy := range s.strategies {
		hints = append(hints, strategy.Find(annotations)...)
	}
	return hints
}
