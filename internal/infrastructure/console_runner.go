package infrastructure

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/beelzebufo98/Hangman/internal/application"
	"github.com/beelzebufo98/Hangman/internal/domain"
)

type ConsoleRunner struct {
	reader  *bufio.Reader
	usecase *application.GameUseCase
}

func NewConsoleRunner(usecase *application.GameUseCase) *ConsoleRunner {
	return &ConsoleRunner{
		reader:  bufio.NewReader(os.Stdin),
		usecase: usecase,
	}
}

func (r *ConsoleRunner) Run() {
	fmt.Println("=== Виселица ===")
	cats := r.usecase.Categories()
	fmt.Println("\nВыберите категорию (Enter — случайная):")
	printNumbered(cats)
	cat := readMenuChoice(r.reader, cats)

	levelLabels := r.usecase.Levels(cat)

	fmt.Println("\nВыберите уровень сложности (Enter — случайный):")
	printNumbered(levelLabels)
	lvl := parseDifficulty(readMenuChoice(r.reader, levelLabels))

	category, level, game, hint := r.usecase.NewGame(cat, lvl, nil)

	drawer := NewAsciiDrawer(level)
	game.Session.MaxAttempts = drawer.MaxStages()

	fmt.Printf("\nКатегория: %s | Сложность: %s\n", category, level.String())
	fmt.Printf("Допустимых ошибок: %d\n", game.Session.MaxAttempts)
	fmt.Println("(Подсказка доступна по запросу после промаха)")

	hintShown := false

	for {
		fmt.Println()
		fmt.Printf("Слово: %s\n", game.RevealedWord())

		if guessed := r.usecase.GuessedLetters(*game.Session); guessed != "" {
			fmt.Printf("Введённые буквы: %s\n", guessed)
		}

		if game.Finished() {
			if game.Status == domain.Win {
				fmt.Printf("\nПобеда! Слово: %q\n", string(game.Session.Secret))
			} else {
				fmt.Printf("\nПоражение. Слово было: %q\n", string(game.Session.Secret))
			}
			return
		}

		rn, ok := readSingleLetter(r.reader)
		if !ok {
			continue
		}

		hit, repeated := game.GuessLetter(rn)
		switch {
		case repeated:
			fmt.Println("Эта буква уже вводилась, попробуйте другую.")
		case hit:
			fmt.Println("Есть совпадение!")
		default:
			fmt.Println("Промах.")
			if !hintShown && askYesNo(r.reader, "Показать подсказку? [y/N]: ") {
				fmt.Printf("Подсказка: %s\n", hint)
				hintShown = true
			}
		}

		left := game.Session.MaxAttempts - game.Session.Attempts
		fmt.Printf("Осталось попыток: %d\n", left)
	}
}

func readSingleLetter(in *bufio.Reader) (rune, bool) {
	fmt.Print("\nВведите букву: ")
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)
	if line == "" {
		return 0, false
	}
	if utf8.RuneCountInString(line) != 1 {
		fmt.Println("Введите ровно одну букву.")
		return 0, false
	}
	r, _ := utf8.DecodeRuneInString(line)
	return unicode.ToLower(r), true
}

func printNumbered(items []string) {
	for i, v := range items {
		fmt.Printf("  %d) %s\n", i+1, v)
	}
}

func readMenuChoice(in *bufio.Reader, items []string) string {
	fmt.Print("> ")
	raw, _ := in.ReadString('\n')
	raw = strings.TrimSpace(raw)
	for _, v := range items {
		if strings.EqualFold(v, raw) {
			return v
		}
	}
	return ""
}

func askYesNo(in *bufio.Reader, prompt string) bool {
	fmt.Print(prompt)
	ans, _ := in.ReadString('\n')
	ans = strings.TrimSpace(strings.ToLower(ans))
	return ans == "y" || ans == "yes" || ans == "д" || ans == "да"
}

func parseDifficulty(s string) domain.Difficulty {
	s = strings.ToLower(strings.TrimSpace(s))
	switch s {
	case strings.ToLower(domain.Medium.String()):
		return domain.Medium
	case strings.ToLower(domain.Hard.String()):
		return domain.Hard
	case strings.ToLower(domain.Easy.String()):
		return domain.Easy
	default:
		return 0
	}
}
