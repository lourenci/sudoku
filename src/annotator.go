package board

type Annotations map[int]map[int][]int

func NewAnnotation() Annotations {
	return make(Annotations)
}

func (a Annotations) Annotate(b Board) Annotations {
	for i := 0; i <= 8; i++ {
		a[i] = make(map[int][]int)

		for j := 0; j <= 8; j++ {
			if b[i][j] == 0 {
				a[i][j] = b.findMissingNumbersInCell(i, j)
			}
		}
	}
	return a
}
