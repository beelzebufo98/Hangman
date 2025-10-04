package domain

import "testing"

func TestNewSession_InitialState(t *testing.T) {
	s := NewSession("кот", 5)

	if string(s.Secret) != "кот" {
		t.Fatalf("secret mismatch: %q", string(s.Secret))
	}
	if s.MaxAttempts != 5 || s.Attempts != 0 {
		t.Fatalf("attempts: got %d/%d", s.Attempts, s.MaxAttempts)
	}
	if s.Won || s.Lost {
		t.Fatalf("won/lost must be false at start")
	}
	if len(s.Revealed) != len(s.Secret) {
		t.Fatalf("revealed len mismatch")
	}
	for i, r := range s.Revealed {
		if r != '_' {
			t.Fatalf("revealed[%d]=%q want '_'", i, r)
		}
	}
}
