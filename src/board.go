package board

type Board [9][9]int
type Annotations map[int]map[int][]int

func (b Board) Annotate() Annotations {
	annotations := make(Annotations)
	missingNumbersInRow := make(map[int][]int)
	missingNumbersInCol := make(map[int][]int)
	missingNumbersInHouse := make(map[int][]int)

	for i, row := range b {
		missingNumbersInRow[i] = FindMissingNumbers(row[:])
	}

	for col := 0; col <= 8; col++ {
		var column [9]int

		for i, row := range b {
			column[i] = row[col]
		}
		missingNumbersInCol[col] = FindMissingNumbers(column[:])
	}

	for house := 0; house <= 8; house++ {
		houseNumbers := b.getHouse(house)
		missingNumbersInHouse[house] = FindMissingNumbers(houseNumbers[:])
	}

	for i := 0; i <= 8; i++ {
		annotations[i] = make(map[int][]int)

		for j := 0; j <= 8; j++ {
			if b[i][j] == 0 {
				var isNumberMissing bool

				for number := 1; number <= 9; number++ {
					isNumberMissing = FindIndex(missingNumbersInRow[i], number) != nil
					if !isNumberMissing {
						continue
					}

					isNumberMissing = FindIndex(missingNumbersInCol[j], number) != nil
					if !isNumberMissing {
						continue
					}

					isNumberMissing = FindIndex(missingNumbersInHouse[GetHouseOfACell(i, j)], number) != nil
					if !isNumberMissing {
						continue
					}

					annotations[i][j] = append(annotations[i][j], number)
				}
			}
		}
	}
	return annotations
}

func (b Board) GetHouses() map[int][9]int {
	housesNumbers := make(map[int][9]int)

	housesNumbers[0] = b.getHouse(0)
	housesNumbers[1] = b.getHouse(1)
	housesNumbers[2] = b.getHouse(2)
	housesNumbers[3] = b.getHouse(3)
	housesNumbers[4] = b.getHouse(4)
	housesNumbers[5] = b.getHouse(5)
	housesNumbers[6] = b.getHouse(6)
	housesNumbers[7] = b.getHouse(7)
	housesNumbers[8] = b.getHouse(8)

	return housesNumbers
}

func (b Board) getHouse(houseNumber int) [9]int {
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

func GetHouseOfACell(rowNumber, colNumber int) int {
	for i := 0; i <= 8; i++ {
		startRow := i / 3 * 3
		startCol := i % 3 * 3

		if startRow+3 > rowNumber && startCol+3 > colNumber {
			return i
		}
	}

	panic("Out of range")
}
