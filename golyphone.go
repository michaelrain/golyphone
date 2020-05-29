package golyphone

import (
	"strings"
)

// Polyphone return the same value for words with same speech
func Polyphone(s string) (string, error) {
	s = strings.ToUpper(s)
	s = latokir(s)
	s = latclear(s)

	return s, nil
}
