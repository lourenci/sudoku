package board

func AnnotateMissingNumbers(l []int) map[int][]int {
	annotations := make(map[int][]int)
	const NOT_FILLED = 0

	for i, cell := range l {
		if cell == NOT_FILLED {
			annotations[i] = FindMissingNumbers(l)
		}
	}

	return annotations
}

func FindMissingNumbers(filledNumbers []int) []int {
	var missingNumbers []int

	for number := 1; number <= 9; number++ {
		if !isNumberInLine(filledNumbers, number) {
			missingNumbers = append(missingNumbers, number)
		}
	}

	return missingNumbers
}

func isNumberInLine(l []int, number int) bool {
	for _, cell := range l {
		if cell == number {
			return true
		}
	}

	return false
}
