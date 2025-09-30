package application

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"unicode"

	"github.com/beelzebufo98/Hangman/internal/domain"
	"github.com/beelzebufo98/Hangman/pkg/word"
)

type GameService struct{ drawer *HangmanEngine }

func NewGameService() *GameService { return &GameService{drawer: Engine()} }

func (s *GameService) Evaluate(secret, guess string) (domain.Result, error) {
	if !word.SameRuneLen(secret, guess) {
		return domain.Result{}, fmt.Errorf("слова должны быть одинаковой длины")
	}
	masked := word.MaskEqualPositions(secret, guess)
	status := domain.StatusNEG
	if masked == secret {
		status = domain.StatusPOS
	}
	return domain.Result{Masked: masked, Status: status}, nil
}

func (s *GameService) NewRandomGame() (category, level string, secret string, hint string, sess domain.Session) {
	cat := keys(Words)[randInt(len(Words))]
	lvl := keys(Words[cat])[randInt(len(Words[cat]))]
	w := Words[cat][lvl][randInt(len(Words[cat][lvl]))]

	max := s.drawer.MaxStages()
	sess = domain.NewSession(w.Text, max)
	return cat, lvl, w.Text, w.Hint, sess
}

func (s *GameService) ApplyGuess(sess *domain.Session, r rune) {
	r = unicode.ToLower(r)
	if sess.Won || sess.Lost {
		return
	}
	if sess.Guessed[r] {
		return
	}
	sess.Guessed[r] = true

	hit := false
	for i, ch := range sess.Secret {
		if unicode.ToLower(ch) == r {
			sess.Revealed[i] = sess.Secret[i]
			hit = true
		}
	}
	if !hit {
		sess.Attempts++
	}
	sess.Won = string(sess.Revealed) == string(sess.Secret)
	sess.Lost = sess.Attempts >= sess.MaxAttempts
}

func (s *GameService) Render(sess domain.Session) string {
	return s.drawer.Stage(sess.Attempts) + "\n" + string(sess.Revealed)
}

func randInt(n int) int {
	if n <= 0 {
		return 0
	}
	m, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
	return int(m.Int64())
}

func keys[M ~map[string]V, V any](m M) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
