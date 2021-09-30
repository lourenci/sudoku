package strategies

import "github.com/lourenci/sudoku/solver"

type NakedSingle struct{}

func (n NakedSingle) Find(annotations solver.Annotations) []solver.Coordinate {
	var findedNumbers []solver.Coordinate

	for x, col := range annotations {
		for y, numbers := range col {
			if len(numbers) == 1 {
				findedNumbers = append(findedNumbers, solver.Coordinate{x, y, numbers[0]})
			}
		}
	}

	return findedNumbers
}
