package sudoku

type Coordinate struct {
	x, y int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x, y}
}
