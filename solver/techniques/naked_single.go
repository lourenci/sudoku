package techniques

type NakedSingle struct{}

func (NakedSingle) Name() string {
	return "Naked Single"
}

func (NakedSingle) Solve(board Board) {
	for c, annotations := range board.Annotations() {
		if len(annotations) == 1 {
			board.Fill(c, annotations[0])
		}
	}
}
