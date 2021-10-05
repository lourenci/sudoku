package strategies

import "github.com/lourenci/sudoku/solver"

type NakedPair struct {}

func (NakedPair) Find(annotations solver.Annotations) []solver.AnnotationHint {
	return []solver.AnnotationHint{{}}
}
