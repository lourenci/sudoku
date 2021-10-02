package strategies

import (
	"github.com/lourenci/sudoku"
	"github.com/lourenci/sudoku/solver"
)

type HiddenSingle struct{}

func (h HiddenSingle) Find(annotations solver.Annotations) []solver.Hint {
	var findedNumbers []solver.Hint

	for x, col := range annotations {
		for y, numbers := range col {
			if len(numbers) == 1 {
				coordinate := sudoku.Coordinate{X: x, Y: y}
				findedNumbers = append(findedNumbers, solver.Hint{coordinate, numbers[0]})
			}
		}
	}

	return findedNumbers
}
