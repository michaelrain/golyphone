package golyphone

import (
	"regexp"
)

func latclear(s string) string {
	re := regexp.MustCompile("[^а-яА-Я]")
	r := re.ReplaceAll([]byte(s), []byte(""))
	return string(r)
}
