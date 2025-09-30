package main

import (
	"fmt"
	"os"

	"github.com/beelzebufo98/Hangman/internal/application"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Предполагалось использование двух аргументов: <загаданное_слово> <проверочное_слово>")
		os.Exit(1)
	}

	secret, guess := os.Args[1], os.Args[2]
	svc := application.NewGameService()
	res, err := svc.Evaluate(secret, guess)
	if err != nil {
		lr := func(s string) int { return len([]rune(s)) }
		fmt.Fprintf(
			os.Stderr,
			"Ошибка: %v (загаданное=%q, длина=%d; проверочное=%q, длина=%d)\n",
			err, secret, lr(secret), guess, lr(guess),
		)
		os.Exit(1)
	}

	fmt.Printf("%s;%s\n", res.Masked, res.Status)
}
