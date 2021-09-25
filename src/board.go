package board

type Board [9][9]int

func (b Board) getNumbersOfCol(colNumber int) [9]int {
	var n [9]int

	for i, row := range b {
		n[i] = row[colNumber]
	}

	return n
}

func (b Board) getNumbersOfHouse(houseNumber int) [9]int {
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

func GetHouseNumberOfCell(rowNumber, colNumber int) int {
	for i := 0; i <= 8; i++ {
		startRow := i / 3 * 3
		startCol := i % 3 * 3

		if startRow+3 > rowNumber && startCol+3 > colNumber {
			return i
		}
	}

	panic("Out of range")
}

func (b Board) findMissingNumbersInCell(i int, j int) []int {
	var missingNumbers []int

	var isNumberMissing bool

	for number := 1; number <= 9; number++ {
		isNumberMissing = FindIndex(b[i][:], number) == nil
		if !isNumberMissing {
			continue
		}

		colNumbers := b.getNumbersOfCol(j)
		isNumberMissing = FindIndex(colNumbers[:], number) == nil
		if !isNumberMissing {
			continue
		}

		numbersInHouse := b.getNumbersOfHouse(GetHouseNumberOfCell(i, j))
		isNumberMissing = FindIndex(numbersInHouse[:], number) == nil
		if !isNumberMissing {
			continue
		}

		missingNumbers = append(missingNumbers, number)
	}

	return missingNumbers
}
