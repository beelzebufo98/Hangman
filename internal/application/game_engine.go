package application

type HangmanEngine struct{ stages []string }

func NewHangmanEngine(level string) *HangmanEngine {
	switch level {
	case "Сложный":
		return engineHard()
	case "Средний":
		return engineMedium()
	default:
		return engineEasy()
	}
}

func (h *HangmanEngine) Stage(errors int) string {
	if errors < 0 {
		errors = 0
	}
	if errors >= len(h.stages) {
		errors = len(h.stages) - 1
	}
	return h.stages[errors]
}

func (h *HangmanEngine) MaxStages() int { return len(h.stages) - 1 }

func engineEasy() *HangmanEngine {
	s := []string{
		`
 ------
 |    |
 |
 |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |    |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   /|\
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   /|\
 |   /
 |
=========
`,
		`
 ------
 |    |
 |    X
 |   /|\
 |   / \
 |
=========
`,
	}
	return &HangmanEngine{stages: s}
}

func engineMedium() *HangmanEngine {
	s := []string{
		`
 ------
 |    |
 |
 |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |    |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   /|
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   /|\
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   /|\
 |   /
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   /|\
 |   / \
 |
=========
`,
		`
 ------
 |    |
 |    X
 |   /|\
 |   / \
 |
=========
`,
	}
	return &HangmanEngine{stages: s}
}

func engineHard() *HangmanEngine {
	s := []string{
		`
 ------
 |    |
 |
 |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |    |
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   \|
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   \|/
 |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   \|/
 |    |
 |
=========
`,
		`
 ------
 |    |
 |    O
 |   \|/
 |    |
 |   /
=========
`,
		`
 ------
 |    |
 |    O
 |   \|/
 |    |
 |   / \
=========
`,
		`
 ------
 |    |
 |   \O
 |   \|/
 |    |
 |   / \
=========
`,
		`
 ------
 |    |
 |   \O/
 |   \|/
 |    |
 |   / \
=========
`,
		`
 ------
 |    |
 |    X
 |   /|\
 |    |
 |   / \
=========
`,
	}
	return &HangmanEngine{stages: s}
}
