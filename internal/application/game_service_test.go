package application

import (
	"testing"
	"unicode"

	"github.com/beelzebufo98/Hangman/internal/domain"
	"github.com/beelzebufo98/Hangman/pkg/word"
)

func newSess(secret string, max int) domain.Session {
	return domain.NewSession(secret, max)
}

func TestEvaluate_CaseInsensitive_POS(t *testing.T) {
	svc := NewGameService()

	res, err := svc.Evaluate("волокно", "ВОЛОКНО")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Status != domain.StatusPOS {
		t.Fatalf("want POS, got %s", res.Status)
	}
}

func TestApplyGuess_Hit_RevealsAllOccurrences_NoAttemptIncrement(t *testing.T) {
	svc := NewGameService()
	sess := newSess("АрбаЛет", 6)

	out := svc.ApplyGuess(&sess, 'А')
	if !out.Hit {
		t.Fatalf("expected hit")
	}
	if sess.Attempts != 0 {
		t.Fatalf("attempts incremented on hit: %d", sess.Attempts)
	}

	for i, r := range sess.Secret {
		if unicode.ToLower(r) == 'a' {
			if sess.Revealed[i] != r {
				t.Fatalf("position %d not revealed correctly: got %q want %q", i, sess.Revealed[i], r)
			}
		}
	}
}

func TestApplyGuess_Miss_IncrementsAttempts_AndLossOnMax(t *testing.T) {
	svc := NewGameService()
	sess := newSess("кот", 1)
	out := svc.ApplyGuess(&sess, 'x')
	if out.Hit {
		t.Fatalf("expected miss, got hit")
	}
	if sess.Attempts != 1 {
		t.Fatalf("want attempts=1, got %d", sess.Attempts)
	}
	if !sess.Lost {
		t.Fatalf("expected Lost=true on reaching max attempts")
	}
	if sess.Won {
		t.Fatalf("won must be false")
	}
}

func TestApplyGuess_Win_SetsWon(t *testing.T) {
	svc := NewGameService()
	sess := newSess("кот", 6)

	svc.ApplyGuess(&sess, 'К')
	if sess.Won || sess.Lost {
		t.Fatalf("should not be final yet")
	}
	svc.ApplyGuess(&sess, 'о')
	if sess.Won || sess.Lost {
		t.Fatalf("should not be final yet")
	}
	svc.ApplyGuess(&sess, 'т')
	if !sess.Won {
		t.Fatalf("expected Won=true after revealing all letters")
	}
	if sess.Lost {
		t.Fatalf("Lost must be false on win")
	}
	if sess.Attempts != 0 {
		t.Fatalf("no mistakes were made, attempts must be 0, got %d", sess.Attempts)
	}
	if string(sess.Revealed) != string(sess.Secret) {
		t.Fatalf("revealed must equal secret: %q vs %q", string(sess.Revealed), string(sess.Secret))
	}
}

func TestApplyGuess_Repeated_DoesNotChangeState(t *testing.T) {
	svc := NewGameService()
	sess := newSess("море", 6)

	first := svc.ApplyGuess(&sess, 'о')
	if !first.Hit || first.Repeated {
		t.Fatalf("first guess should be a fresh hit")
	}

	secret := string(sess.Secret)
	revealed := string(sess.Revealed)
	attempts := sess.Attempts

	second := svc.ApplyGuess(&sess, 'О')
	if !second.Repeated {
		t.Fatalf("expected Repeated=true on second same guess")
	}
	if secret != string(sess.Secret) || revealed != string(sess.Revealed) || attempts != sess.Attempts {
		t.Fatalf("state changed on repeated guess")
	}
}

func TestEvaluate_NEG_WithMasked(t *testing.T) {
	svc := NewGameService()

	secret := "абвг"
	guess := "АбДЕ"
	res, err := svc.Evaluate(secret, guess)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantMask := word.MaskEqualPositions(secret, guess)
	if res.Masked != wantMask {
		t.Fatalf("masked mismatch: got %q want %q", res.Masked, wantMask)
	}
	if res.Status != domain.StatusNEG {
		t.Fatalf("want NEG, got %s", res.Status)
	}
}

func TestEvaluate_LengthMismatch_Error(t *testing.T) {
	svc := NewGameService()
	if _, err := svc.Evaluate("кот", "коты"); err == nil {
		t.Fatalf("expected error on length mismatch")
	}
}

func TestApplyGuess_RepeatAfterMiss_NoExtraAttempt(t *testing.T) {
	svc := NewGameService()
	sess := domain.NewSession("мир", 6)

	o1 := svc.ApplyGuess(&sess, 'x')
	if o1.Hit || o1.Repeated {
		t.Fatalf("first x must be a miss")
	}
	if sess.Attempts != 1 {
		t.Fatalf("attempts must be 1 after first miss, got %d", sess.Attempts)
	}
	o2 := svc.ApplyGuess(&sess, 'X')
	if !o2.Repeated {
		t.Fatalf("second same miss must be repeated")
	}
	if sess.Attempts != 1 {
		t.Fatalf("attempts changed on repeated miss: %d", sess.Attempts)
	}
}

func TestWinAndLossPaths(t *testing.T) {
	svc := NewGameService()

	s1 := domain.NewSession("дом", 5)
	svc.ApplyGuess(&s1, 'д')
	svc.ApplyGuess(&s1, 'О')
	res := svc.ApplyGuess(&s1, 'м')
	if !res.Hit || !s1.Won || s1.Lost {
		t.Fatalf("expected win state")
	}
	if string(s1.Revealed) != string(s1.Secret) {
		t.Fatalf("revealed not equal secret on win")
	}

	s2 := domain.NewSession("кот", 2)
	svc.ApplyGuess(&s2, 'x')
	svc.ApplyGuess(&s2, 'y')
	if !s2.Lost || s2.Won {
		t.Fatalf("expected loss state after reaching max attempts")
	}
}

func TestLevels_OrderIsAscending(t *testing.T) {
	svc := NewGameService()
	cats := svc.Categories()
	if len(cats) == 0 {
		t.Fatalf("no categories")
	}
	levels := svc.Levels(cats[0])
	if len(levels) < 3 {
		t.Fatalf("expected >=3 levels, got %v", levels)
	}
	want := []string{"Лёгкий", "Средний", "Сложный"}
	for i := range want {
		if levels[i] != want[i] {
			t.Fatalf("levels order unexpected: %v", levels)
		}
	}
}

func TestNewGame_ValidOutputs(t *testing.T) {
	svc := NewGameService()
	cat, lvl, wordStr, hint, sess := svc.NewGame("", "")
	if cat == "" || lvl == "" {
		t.Fatalf("category/level must be chosen")
	}
	if len(wordStr.Text) == 0 || len(sess.Secret) == 0 {
		t.Fatalf("word/secret must be non-empty")
	}
	if len(sess.Revealed) != len(sess.Secret) {
		t.Fatalf("revealed len must equal secret len")
	}
	if hint == "" {
		t.Fatalf("hint should be present for bonus feature")
	}
}
