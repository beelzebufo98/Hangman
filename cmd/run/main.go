package main

import (
	"fmt"
	"os"

	"github.com/beelzebufo98/Hangman/internal/application"
	"github.com/beelzebufo98/Hangman/internal/domain"
	"github.com/beelzebufo98/Hangman/internal/infrastructure"
)

func main() {
	switch len(os.Args) {
	case 1:
		provider := infrastructure.NewStaticWordProvider()
		drawer := infrastructure.NewAsciiDrawer(domain.Easy)
		repo := infrastructure.NewMemorySessionRepo()

		usecase := application.NewGameUseCase(provider, drawer, repo)
		runner := infrastructure.NewConsoleRunner(usecase)
		runner.Run()

	case 3:
		secret, guess := os.Args[1], os.Args[2]
		usecase := application.NewEvalOnlyGameUseCase()
		res, err := usecase.Evaluate(secret, guess)
		if err != nil {
			lr := func(s string) int { return len([]rune(s)) }
			fmt.Fprintf(os.Stderr,
				"Ошибка: %v (загаданное=%q, длина=%d; проверочное=%q, длина=%d)\n",
				err, secret, lr(secret), guess, lr(guess),
			)
			os.Exit(1)
		}
		fmt.Printf("%s;%s\n", res.Masked, res.Status)

	default:
		fmt.Fprintln(os.Stderr,
			"Использование:\n"+
				"  ./hangman                # интерактивный режим\n"+
				"  ./hangman <слово> <попытка>  # неинтерактивный режим",
		)
		os.Exit(1)
	}
}
