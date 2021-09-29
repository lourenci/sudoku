package solver

type HiddenSingle struct {}

func (h HiddenSingle) Find(annotations Annotations) []Coordinate {
	var findedNumbers []Coordinate

	for x, col := range annotations {
		for y, numbers := range col {
			if len(numbers) == 1 {
				findedNumbers = append(findedNumbers, Coordinate{x, y, numbers[0]})
			}
		}
	}

	return findedNumbers
}
