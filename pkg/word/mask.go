package word

import "strings"

func MaskEqualPositions(secret, guess string) string {
	s := []rune(secret)
	ls := []rune(strings.ToLower(secret))
	lg := []rune(strings.ToLower(guess))

	res := make([]rune, len(s))
	for i := range s {
		if ls[i] == lg[i] {
			res[i] = s[i]
		} else {
			res[i] = '*'
		}
	}
	return string(res)
}

func SameRuneLen(a, b string) bool { return len([]rune(a)) == len([]rune(b)) }
