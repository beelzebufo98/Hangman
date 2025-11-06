package domain

type Session struct {
	Secret      []rune
	Revealed    []rune
	Attempts    int
	MaxAttempts int
	Guessed     map[rune]bool
	Won         bool
	Lost        bool
}

func NewSession(secret string, max int) Session {
	s := []rune(secret)
	r := make([]rune, len(s))
	for i := range r {
		r[i] = '_'
	}
	return Session{
		Secret:      s,
		Revealed:    r,
		Attempts:    0,
		MaxAttempts: max,
		Guessed:     map[rune]bool{},
	}
}
