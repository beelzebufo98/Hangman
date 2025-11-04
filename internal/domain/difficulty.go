package domain

type Difficulty int

const (
	Easy Difficulty = iota + 1
	Medium
	Hard
)

func (d Difficulty) String() string {
	switch d {
	case Easy:
		return "Лёгкий"
	case Medium:
		return "Средний"
	case Hard:
		return "Сложный"
	default:
		return "Не назначен"
	}
}
