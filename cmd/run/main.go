package main

import (
	"fmt"
	"os"

	"github.com/beelzebufo98/Hangman/internal/application"
	"github.com/beelzebufo98/Hangman/internal/infrastructure"
)

func main() {
	switch len(os.Args) {
	case 3:
		// неинтерактив: run <secret> <guess>
		secret, guess := os.Args[1], os.Args[2]
		svc := application.NewGameService()
		res, err := svc.Evaluate(secret, guess)
		if err != nil {
			lr := func(s string) int { return len([]rune(s)) }
			fmt.Fprintf(os.Stderr, "Ошибка: %v (загаданное=%q, длина=%d; проверочное=%q, длина=%d)\n",
				err, secret, lr(secret), guess, lr(guess))
			os.Exit(1)
		}
		fmt.Printf("%s;%s\n", res.Masked, res.Status)
	default:
		// интерактив без аргументов (или с мусорными — тоже сюда)
		infrastructure.RunInteractive()
	}
}
