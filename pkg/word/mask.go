package word

func MaskEqualPositions(secret, guess string) string {
	s := []rune(secret)
	g := []rune(guess)
	res := make([]rune, len(s))
	for i := range s {
		if s[i] == g[i] {
			res[i] = s[i]
		} else {
			res[i] = '*'
		}
	}
	return string(res)
}

func SameRuneLen(a, b string) bool { return len([]rune(a)) == len([]rune(b)) }
