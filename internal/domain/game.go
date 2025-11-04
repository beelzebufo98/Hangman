package domain

import "unicode"

type Game struct {
	Session *Session
	Status  GameStatus
}

func NewGame(secret string, maxAttempts int) *Game {
	sess := NewSession(secret, maxAttempts)
	return &Game{
		Session: &sess,
		Status:  InProgress,
	}
}

func (g *Game) GuessLetter(r rune) (hit bool, repeated bool) {
	s := g.Session
	r = unicode.ToLower(r)

	if s.Won || s.Lost {
		return false, false
	}
	if s.Guessed[r] {
		return false, true
	}

	s.Guessed[r] = true
	hit = false

	for i, ch := range s.Secret {
		if unicode.ToLower(ch) == r {
			s.Revealed[i] = s.Secret[i]
			hit = true
		}
	}

	if !hit {
		s.Attempts++
	}

	s.Won = string(s.Revealed) == string(s.Secret)
	s.Lost = s.Attempts >= s.MaxAttempts

	switch {
	case s.Won:
		g.Status = Win
	case s.Lost:
		g.Status = Lose
	default:
		g.Status = InProgress
	}

	return hit, false
}

func (g *Game) RevealedWord() string {
	return string(g.Session.Revealed)
}

func (g *Game) Finished() bool {
	return g.Status == Win || g.Status == Lose
}
