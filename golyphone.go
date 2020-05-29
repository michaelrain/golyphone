package golyphone

import (
	"regexp"
	"strings"
)

// Polyphone return the same value for words with same speech
func Polyphone(s string) (string, error) {
	s = strings.ToUpper(s)
	s = latokir(s)

	re := regexp.MustCompile("[^а-яА-Я]")

	return s, nil
}
