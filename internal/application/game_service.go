package application

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"unicode"

	"github.com/beelzebufo98/Hangman/internal/domain"
	"github.com/beelzebufo98/Hangman/pkg/word"
)

type GameService struct{ drawer *HangmanEngine }

func NewGameService() *GameService { return &GameService{drawer: NewHangmanEngine("Лёгкий")} }

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

func (s *GameService) Categories() []string {
	ks := mapKeys(Words)
	sort.Strings(ks)
	return ks
}
func (s *GameService) Levels(cat string) []string {
	cat = s.resolveCategory(cat)
	ks := mapKeys(Words[cat])
	sort.Strings(ks)
	return ks
}

func (s *GameService) NewGame(cat, lvl string) (category, level string, w Word, hint string, sess domain.Session) {
	category = s.resolveCategory(cat)
	level = s.resolveLevel(category, lvl)

	pool := Words[category][level]
	w = pool[randInt(len(pool))]

	s.drawer = NewHangmanEngine(level)
	max := s.drawer.MaxStages()

	sess = domain.NewSession(w.Text, max)
	return category, level, w, w.Hint, sess
}

type GuessOutcome struct {
	Repeated bool
	Hit      bool
}

func (s *GameService) ApplyGuess(sess *domain.Session, r rune) GuessOutcome {
	r = unicode.ToLower(r)

	if sess.Won || sess.Lost {
		return GuessOutcome{}
	}
	if sess.Guessed[r] {
		return GuessOutcome{Repeated: true}
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
	return GuessOutcome{Hit: hit}
}

func (s *GameService) GuessedLetters(sess domain.Session) string {
	if len(sess.Guessed) == 0 {
		return ""
	}
	letters := make([]string, 0, len(sess.Guessed))
	for r := range sess.Guessed {
		letters = append(letters, string(r))
	}
	sort.Strings(letters)
	return strings.Join(letters, " ")
}

func (s *GameService) Render(sess domain.Session) string {
	return s.drawer.Stage(sess.Attempts) + "\n" + string(sess.Revealed)
}

func (s *GameService) resolveCategory(in string) string {
	if in == "" {
		ks := mapKeys(Words)
		return ks[randInt(len(ks))]
	}
	for k := range Words {
		if strings.EqualFold(k, in) {
			return k
		}
	}
	ks := mapKeys(Words)
	return ks[randInt(len(ks))]
}
func (s *GameService) resolveLevel(cat, lvl string) string {
	if lvl == "" {
		ks := mapKeys(Words[cat])
		return ks[randInt(len(ks))]
	}
	for k := range Words[cat] {
		if strings.EqualFold(k, lvl) {
			return k
		}
	}
	ks := mapKeys(Words[cat])
	return ks[randInt(len(ks))]
}
func mapKeys[M ~map[string]V, V any](m M) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
func randInt(n int) int {
	if n <= 0 {
		return 0
	}
	m, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
	return int(m.Int64())
}
