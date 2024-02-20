package polyphone

import (
	"bytes"
)

func deldoubles(s string) string {
	var buf bytes.Buffer
	var last rune

	for i, r := range s {
		if r != last || i == 0 {
			buf.WriteRune(r)
			last = r
		}
	}

	return buf.String()
}
