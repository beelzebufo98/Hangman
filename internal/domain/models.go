package domain

type Status string

const (
	StatusPOS Status = "POS"
	StatusNEG Status = "NEG"
)

type Result struct {
	Masked string
	Status Status
}
