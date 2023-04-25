package techniques

type FullHouse struct{}

func (FullHouse) Name() string {
	return "Full House"
}

func (FullHouse) Solve(board Board) {
	for c, annotations := range board.Annotations() {
		if len(annotations) == 1 {
			board.Fill(c, annotations[0])
		}
	}
}
