package sudoku

import (
	"strconv"

	"github.com/lourenci/sudoku/src/utils"
)

type Board [9][9]int

// TODO should this be moved to a parser?
// TODO validate the board
func NewBoardFromString(board string) Board {
	var b Board

	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			number, _ := strconv.Atoi(string(board[i * 9 + j]))
			b[i][j] = number
		}
	}

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

func (b Board) MissingNumbersInCell(rowNumber int, colNumber int) []int {
	var missingNumbers []int

	if b.isCellFilled(rowNumber, colNumber) {
		return nil
	}

	for number := 1; number <= 9; number++ {
		if !b.isNumberFilledInBoard(rowNumber, colNumber, number) {
			missingNumbers = append(missingNumbers, number)
		}
	}

	return missingNumbers
}

func (b Board) isCellFilled(rowNumber int, colNumber int) bool {
	if b[rowNumber][colNumber] != 0 {
		return true
	}
	return false
}

func (b Board) isNumberFilledInBoard(rowNumber int, colNumber int, number int) bool {
	if b.isNumberFilledInRow(rowNumber, number) {
		return true
	}

	if b.isNumberFilledInCol(colNumber, number) {
		return true
	}

	if b.isNumberFilledInHouse(getHouseNumberOfCell(rowNumber, colNumber), number) {
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

func getHouseNumberOfCell(rowNumber, colNumber int) int {
	for i := 0; i <= 8; i++ {
		startRow := i / 3 * 3
		startCol := i % 3 * 3

		if startRow+3 > rowNumber && startCol+3 > colNumber {
			return i
		}
	}

	panic("Out of range")
}

func (b Board) getFilledNumbersOfHouse(houseNumber int) [9]int {
	startRow := houseNumber / 3 * 3
	endRow := startRow + 3
	startCol := houseNumber % 3 * 3
	endCol := startCol + 3

	var numbers [9]int
	qtyNumbers := 0
	for i := startRow; i < endRow; i++ {
		for j := startCol; j < endCol; j++ {
			numbers[qtyNumbers] = b[i][j]
			qtyNumbers++
		}
	}

	return numbers
}
