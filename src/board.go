package board

type Board [9][9]int
type Annotations map[int]map[int][]int

func (b Board) Annotate() Annotations {
	annotations := make(Annotations)
	firstRow := b[0]

	missingNumbers := findMissingNumbersInRow(firstRow)
	annotationsOfFirstRow := annotateMissingNumbersInRow(missingNumbers, firstRow)
	annotations[0] = annotationsOfFirstRow

	return annotations
}

func findMissingNumbersInRow(row [9]int) []int {
	var missingNumbers []int

	for number := 1; number <= 9; number++ {
		if !isNumberInRow(number, row) {
			missingNumbers = append(missingNumbers, number)
		}
	}

	return missingNumbers
}

func isNumberInRow(number int, row [9]int) bool {
	for _, cell := range row {
		if cell == number {
			return true
		}
	}

	return false
}

func annotateMissingNumbersInRow(missingNumbers []int, row [9]int) map[int][]int {
	annotations := make(map[int][]int)
	const NOT_FILLED = 0

	for i, cell := range row {
		if cell == NOT_FILLED {
			for _, missingNumber := range missingNumbers {
				annotations[i] = append(annotations[i], missingNumber)
			}
		}
	}

	return annotations
}
