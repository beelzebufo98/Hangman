package word

import "testing"

func TestMaskEqualPositions_CaseInsensitive(t *testing.T) {
	secret := "ВоЛоКнО"
	guess := "волокно"

	m := MaskEqualPositions(secret, guess)
	if m != secret {
		t.Fatalf("mask mismatch: got %q want %q", m, secret)
	}
}

func TestSameRuneLen_UnicodeRunes(t *testing.T) {
	if !SameRuneLen("ёлка", "ЕЛКА") {
		t.Fatalf("expected same rune length for case variants")
	}
	if SameRuneLen("кот", "коты") {
		t.Fatalf("lengths should differ")
	}
}
