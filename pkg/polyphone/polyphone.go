package polyphone

import (
	"strings"
)

// Polyphone return the same value for words with same speech
func Encode(s string) (string, error) {

	s = strings.ToUpper(s)

	s = latokir(s)
	s = latclear(s)
	s = delchars(s)
	s = deldoubles(s)
	s = replaces(s)

	return s, nil
}
