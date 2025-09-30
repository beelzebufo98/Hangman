package application

import (
	"fmt"

	"github.com/beelzebufo98/Hangman/internal/domain"
	"github.com/beelzebufo98/Hangman/pkg/word"
)

type GameService struct{}

func NewGameService() *GameService { return &GameService{} }

func (s *GameService) Evaluate(secret, guess string) (domain.Result, error) {
	if !word.SameRuneLen(secret, guess) {
		return domain.Result{}, fmt.Errorf("length mismatch: words must have equal rune length")
	}
	masked := word.MaskEqualPositions(secret, guess)
	status := domain.StatusNEG
	if masked == secret {
		status = domain.StatusPOS
	}
	return domain.Result{Masked: masked, Status: status}, nil
}
