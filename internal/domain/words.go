package domain

type Word struct {
	Text string
	Hint string
}

type WordProvider interface {
	GetCategories() []string
	GetLevels(category string) []Difficulty
	GetRandomWord(category string, level Difficulty) Word
}
