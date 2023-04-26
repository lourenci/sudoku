package sudoku

import (
	"regexp"
	"strconv"

	"github.com/lourenci/sudoku/modules/collections"
)

type Board struct {
	numbers [9][9]int
}

func NewBoard(numbers [9][9]int) Board {
	return Board{
		numbers: numbers,
	}
}

func ParsedBoard(board string) Board {
	var parsedBoard [9][9]int
	var row [9]int

	board = removeNonDigitsFromString(board)

	if len(board) > 81 {
		panic("please, provide a 9x9 board")
	}

	for i, rune := range board {
		colNumber := i % 9
		if colNumber == 0 {
			row = [9]int{}
		}
		value, _ := strconv.Atoi(string(rune))
		row[colNumber] = value

		rowNumber := i / 9
		parsedBoard[rowNumber] = row
	}

	return NewBoard(parsedBoard)
}

func (r Board) Numbers() [9][9]int {
	return r.numbers
}

func (r Board) Column(number int) []int {
	var column []int

	for _, iv := range r.numbers {
		for j, jv := range iv {
			if j == number && jv != 0 {
				column = append(column, jv)
			}
		}
	}

	return column
}

func (r Board) Row(number int) []int {
	var rowNumbers []int

	for _, v := range r.numbers[number] {
		if v != 0 {
			rowNumbers = append(rowNumbers, v)
		}
	}

	return rowNumbers
}

func (r Board) House(coordinate Coordinate) []int {
	startX := (coordinate.X) / 3 * 3
	endX := startX + 2

	startY := (coordinate.Y) / 3 * 3
	endY := startY + 2

	var rowNumbers []int

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			number := r.numbers[i][j]
			if number != 0 {
				rowNumbers = append(rowNumbers, number)
			}
		}
	}

	return rowNumbers
}

func (r Board) Annotations() map[Coordinate][]int {
	annotations := make(map[Coordinate][]int)
	numbers := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for rowNumber, iv := range r.numbers {
		for columnNumber, cellNumber := range iv {
			if cellNumber == 0 {
				possibleNumbersInRow := collections.Except(
					numbers[:],
					r.Row(rowNumber)[:],
				)
				possibleNumbersInColumn := collections.Except(
					numbers[:],
					r.Column(columnNumber)[:],
				)
				possibleNumbersInHouse := collections.Except(
					numbers[:],
					r.House(NewCoordinate(rowNumber, columnNumber)),
				)
				possibleNumbersInBothRowColAndHouse := collections.Intersect(
					collections.Intersect(possibleNumbersInRow, possibleNumbersInColumn),
					possibleNumbersInHouse,
				)

				annotations[NewCoordinate(rowNumber, columnNumber)] = possibleNumbersInBothRowColAndHouse
			}
		}
	}

	return annotations
}

func (r *Board) Fill(coordinate Coordinate, number int) {
	if number < 1 || number > 9 {
		panic("invalid number")
	}

	r.numbers[coordinate.X][coordinate.Y] = number
}

func removeNonDigitsFromString(str string) string {
	regex := regexp.MustCompile(`\D`)
	return string(regex.ReplaceAll([]byte(str), []byte("")))
}
