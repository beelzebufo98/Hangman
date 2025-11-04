package application

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/beelzebufo98/Hangman/internal/domain"
	"github.com/beelzebufo98/Hangman/pkg/word"
)

type GameService struct {
	drawer domain.Drawer
	words  domain.WordProvider
}

func NewGameService(provider domain.WordProvider, drawer domain.Drawer) *GameService {
	return &GameService{drawer: drawer, words: provider}
}

func NewEvalOnlyService() *GameService { return &GameService{} }

func (s *GameService) Evaluate(secret, guess string) (domain.Result, error) {
	if !word.SameRuneLen(secret, guess) {
		return domain.Result{}, fmt.Errorf("слова должны быть одинаковой длины")
	}
	masked := word.MaskEqualPositions(secret, guess)
	status := domain.InProgress
	if masked == secret {
		status = domain.Win
	}
	return domain.Result{Masked: masked, Status: status}, nil
}

func (s *GameService) Categories() []string {
	return s.words.GetCategories()
}

func (s *GameService) Levels(cat string) []string {
	levels := s.words.GetLevels(cat)
	out := make([]string, len(levels))
	for i, d := range levels {
		out[i] = d.String()
	}
	sort.Slice(out, func(i, j int) bool { return i < j })
	return out
}

func (s *GameService) NewGame(cat string, lvl domain.Difficulty) (category string, level domain.Difficulty, w domain.Word, hint string, sess domain.Session) {
	for _, c := range s.words.GetCategories() {
		if strings.EqualFold(c, cat) {
			category = c
			break
		}
	}
	if category == "" {
		category = s.words.GetCategories()[0]
	}
	level = lvl
	w = s.words.GetRandomWord(category, level)

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
	return s.drawer.Render(sess.Attempts) + "\n" + string(sess.Revealed)
}
