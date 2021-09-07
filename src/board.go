package board

type Board [9][9]int
type Annotations map[int]map[int][]int

func (b Board) Annotate() Annotations {
	return make(Annotations)
}
