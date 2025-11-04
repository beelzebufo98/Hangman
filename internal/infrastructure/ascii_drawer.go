package infrastructure

import "github.com/beelzebufo98/Hangman/internal/domain"

type AsciiDrawer struct{ stages []string }

func NewAsciiDrawer(d domain.Difficulty) *AsciiDrawer {
	switch d {
	case domain.Hard:
		return &AsciiDrawer{stages: stagesHard()}
	case domain.Medium:
		return &AsciiDrawer{stages: stagesMedium()}
	default:
		return &AsciiDrawer{stages: stagesEasy()}
	}
}

func (a *AsciiDrawer) Render(errors int) string {
	if errors < 0 {
		errors = 0
	}
	if errors >= len(a.stages) {
		errors = len(a.stages) - 1
	}
	return a.stages[errors]
}

func (a *AsciiDrawer) MaxStages() int { return len(a.stages) - 1 }

func stagesEasy() []string {
	return []string{
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
}

func stagesMedium() []string {
	return []string{
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
}

func stagesHard() []string {
	return []string{
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
 |    X
 |   /|\
 |    |
 |   / \
=========
`,
	}
}
