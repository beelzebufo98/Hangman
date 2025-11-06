package domain

type GameStatus int

const (
	InProgress GameStatus = iota
	Win
	Lose
)

type Result struct {
	Masked string
	Status GameStatus
}

func (s GameStatus) String() string {
	switch s {
	case Win:
		return "POS"
	case Lose:
		return "NEG"
	default:
		return "NEG"
	}
}
