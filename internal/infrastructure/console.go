package infrastructure

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/beelzebufo98/Hangman/internal/application"
)

func RunInteractive() {
	in := bufio.NewReader(os.Stdin)
	svc := application.NewGameService()

	fmt.Println("=== Виселица ===")
	cats := svc.Categories()
	fmt.Println("\nВыберите категорию (Enter — случайная):")
	printNumbered(cats)
	cat := readMenuChoice(in, cats)

	lvls := svc.Levels(cat)
	fmt.Println("\nВыберите уровень сложности (Enter — случайный):")
	printNumbered(lvls)
	lvl := readMenuChoice(in, lvls)

	category, level, _, hint, sess := svc.NewGame(cat, lvl)

	fmt.Printf("\nКатегория: %s | Сложность: %s\n", category, level)
	fmt.Printf("Допустимых ошибок: %d\n", sess.MaxAttempts)
	fmt.Println("(Подсказка доступна по запросу после промаха)")

	hintShown := false

	for {
		fmt.Println()
		fmt.Println(svc.Render(sess))

		if sess.Won {
			fmt.Printf("\nПобеда! Слово: %q\n", string(sess.Secret))
			return
		}
		if sess.Lost {
			fmt.Printf("\nПоражение. Слово было: %q\n", string(sess.Secret))
			return
		}

		if guessed := svc.GuessedLetters(sess); guessed != "" {
			fmt.Printf("Были буквы: %s\n", guessed)
		}

		fmt.Print("\nВведите букву: ")
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		r, _ := utf8.DecodeRuneInString(line)
		if r == utf8.RuneError {
			continue
		}
		if utf8.RuneCountInString(line) > 1 {
			fmt.Printf("Вы ввели несколько символов, беру первую букву: %q\n", string(r))
		}
		r = unicode.ToLower(r)

		out := svc.ApplyGuess(&sess, r)
		switch {
		case out.Repeated:
			fmt.Println("Эта буква уже вводилась, попробуйте другую.")
		case out.Hit:
			fmt.Println("Есть совпадение!")
		default:
			fmt.Println("Промах.")
			if !hintShown && askYesNo(in, "Показать подсказку? [y/N]: ") {
				fmt.Printf("Подсказка: %s\n", hint)
				hintShown = true
			}
		}

		left := sess.MaxAttempts - sess.Attempts
		fmt.Printf("Осталось попыток: %d\n", left)
	}
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
	if raw == "" {
		return ""
	}
	if idx, err := strconv.Atoi(raw); err == nil && idx >= 1 && idx <= len(items) {
		return items[idx-1]
	}
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
