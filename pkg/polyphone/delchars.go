package polyphone

import (
	"regexp"
)

func delchars(s string) string {
	re := regexp.MustCompile(`[ь,ъ,Ь,Ъ]`)
	r := re.ReplaceAll([]byte(s), []byte(""))
	return string(r)
}
