package sudoku

import (
	"regexp"
	"strconv"
)

type Coordinate struct {
	x, y int
}

type Board struct {
	numbers [9][9]int
}

func NewBoard(numbers [9][9]int) Board {
	return Board{
		numbers: numbers,
	}
}

func Parse(board string) Board {
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
	startX := (coordinate.x) / 3 * 3
	endX := startX + 2

	startY := (coordinate.y) / 3 * 3
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

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x, y}
}

func removeNonDigitsFromString(str string) string {
	regex := regexp.MustCompile(`\D`)
	return string(regex.ReplaceAll([]byte(str), []byte("")))
}
