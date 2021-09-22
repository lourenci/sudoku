package board

type Board [9][9]int
type Annotations map[int]map[int][]int

func (b Board) Annotate() Annotations {
	annotations := make(Annotations)

	for i := 0; i <= 8; i++ {
		annotations[i] = make(map[int][]int)

		for j := 0; j <= 8; j++ {
			if b[i][j] == 0 {
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

					numbersInHouse := b.getHouse(GetHouseOfACell(i, j))
					isNumberMissing = FindIndex(numbersInHouse[:], number) == nil
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

func (b Board) getNumbersOfCol(colNumber int) [9]int {
	var n [9]int

	for i, row := range b {
		n[i] = row[colNumber]
	}

	return n
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
