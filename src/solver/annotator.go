package solver

type Annotations map[int]map[int][]int

func NewAnnotation() Annotations {
	return make(Annotations)
}

type Board interface {
	MissingNumbersInCell(int, int) []int
}

func (a Annotations) Annotate(b Board) Annotations {
	for rowNumber := 0; rowNumber <= 8; rowNumber++ {
		a[rowNumber] = make(map[int][]int)

		for colNumber := 0; colNumber <= 8; colNumber++ {
			a.annotateCellWithMissingNumbers(b, rowNumber, colNumber)
		}
	}
	return a
}

func (a Annotations) annotateCellWithMissingNumbers(b Board, rowNumber int, colNumber int) {
	missingNumbers := b.MissingNumbersInCell(rowNumber, colNumber)
	if missingNumbers != nil {
		a[rowNumber][colNumber] = missingNumbers
	}
}
