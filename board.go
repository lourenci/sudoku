package sudoku

import (
	"github.com/lourenci/sudoku/utils"
)

type Board [9][9]int

type Coordinate struct {
	X int
	Y int
}

func (b Board) Fill(coordinate Coordinate, number int) Board {
	b[coordinate.X][coordinate.Y] = number
	return b
}

func (b Board) IsComplete() bool {
	for _, row := range b {
		for _, number := range row {
			if number == 0 {
				return false
			}
		}
	}

	return true
}

func (b Board) MissingNumbersInCell(coordinate Coordinate) []int {
	var missingNumbers []int

	if b.isCellFilled(coordinate) {
		return nil
	}

	for number := 1; number <= 9; number++ {
		if !b.isNumberFilledInBoard(coordinate, number) {
			missingNumbers = append(missingNumbers, number)
		}
	}

	return missingNumbers
}

func (b Board) isCellFilled(c Coordinate) bool {
	return b[c.X][c.Y] != 0
}

func (b Board) isNumberFilledInBoard(coordinate Coordinate, number int) bool {
	if b.isNumberFilledInRow(coordinate.X, number) {
		return true
	}

	if b.isNumberFilledInCol(coordinate.Y, number) {
		return true
	}

	if b.isNumberFilledInHouse(getHouseNumberOfCell(coordinate), number) {
		return true
	}

	return false
}

func (b Board) isNumberFilledInRow(rowNumber int, number int) bool {
	return utils.FindIndex(b[rowNumber][:], number) != nil
}

func (b Board) isNumberFilledInCol(colNumber int, number int) bool {
	colNumbers := b.getFilledNumbersOfCol(colNumber)
	return utils.FindIndex(colNumbers[:], number) != nil
}

func (b Board) getFilledNumbersOfCol(colNumber int) [9]int {
	var n [9]int

	for i, row := range b {
		n[i] = row[colNumber]
	}

	return n
}

func (b Board) isNumberFilledInHouse(houseNumber int, number int) bool {
	numbersInHouse := b.getFilledNumbersOfHouse(houseNumber)
	return utils.FindIndex(numbersInHouse[:], number) != nil
}

func getHouseNumberOfCell(coordinate Coordinate) int {
	for i := 0; i <= 8; i++ {
		startRow := i / 3 * 3
		startCol := i % 3 * 3

		if startRow+3 > coordinate.X && startCol+3 > coordinate.Y {
			return i
		}
	}

	panic("Out of range")
}

func (b Board) getFilledNumbersOfHouse(houseNumber int) [9]int {
	startCoordinate, endCoordinate := GetCoordinatesOfHouse(houseNumber)

	return b.getFilledNumbersInRange(startCoordinate, endCoordinate)
}

func (b Board) getFilledNumbersInRange(startCoordinate, endCoordinate Coordinate) [9]int {
	var numbers [9]int
	qtyNumbers := 0

	for i := startCoordinate.X; i <= endCoordinate.X; i++ {
		for j := startCoordinate.Y; j <= endCoordinate.Y; j++ {
			numbers[qtyNumbers] = b[i][j]
			qtyNumbers++
		}
	}

	return numbers
}

func GetCoordinatesOfHouse(houseNumber int) (Coordinate, Coordinate) {
	startRow := houseNumber / 3 * 3
	endRow := startRow + 2
	startCol := houseNumber % 3 * 3
	endCol := startCol + 2

	return Coordinate{X: startRow, Y: startCol}, Coordinate{X: endRow, Y: endCol}
}
