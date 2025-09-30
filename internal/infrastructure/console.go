package infrastructure

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/beelzebufo98/Hangman/internal/application"
)

func RunInteractive() {
	svc := application.NewGameService()
	category, level, _, hint, sess := svc.NewRandomGame()

	in := bufio.NewReader(os.Stdin)

	fmt.Printf("Категория: %s | Сложность: %s\n", category, level)
	fmt.Printf("Подсказка: %s\n", hint)
	fmt.Printf("Допустимых ошибок: %d\n\n", sess.MaxAttempts)

	for {
		fmt.Println(svc.Render(sess))
		if sess.Won {
			fmt.Printf("\nУгадали! Слово: %q\n", string(sess.Secret))
			return
		}
		if sess.Lost {
			fmt.Printf("\nПоражение. Слово было: %q\n", string(sess.Secret))
			return
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
		svc.ApplyGuess(&sess, r)
		left := sess.MaxAttempts - sess.Attempts
		fmt.Printf("Осталось попыток: %d\n\n", left)
	}
}
