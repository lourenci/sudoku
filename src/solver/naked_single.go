package solver

type Strategy struct {
	a Annotations
}

type Coordinate struct {
	X      int
	Y      int
	Number int
}

func NewStrategy(a Annotations) Strategy {
	return Strategy{a}
}

func (s Strategy) Find() []Coordinate {
	var findedNumbers []Coordinate

	for x, col := range s.a {
		for y, numbers := range col {
			if len(numbers) == 1 {
				findedNumbers = append(findedNumbers, Coordinate{x, y, numbers[0]})
			}
		}
	}

	return findedNumbers
}
