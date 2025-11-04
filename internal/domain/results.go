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
