package application

type HangmanEngine struct{ stages []string }

func Engine() *HangmanEngine {
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
 |    X
 |   /|\
 |   / \
 |
=========
`,
	}
	return &HangmanEngine{stages: s}
}

func (d *HangmanEngine) Stage(errors int) string {
	if errors < 0 {
		errors = 0
	}
	if errors >= len(d.stages) {
		errors = len(d.stages) - 1
	}
	return d.stages[errors]
}

func (d *HangmanEngine) MaxStages() int { return len(d.stages) - 1 }
