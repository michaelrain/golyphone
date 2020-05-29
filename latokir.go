package golyphone

import "strings"

var substitutes = [8][2]string{
	{"A", "А"},
	{"E", "Е"},
	{"O", "О"},
	{"C", "С"},
	{"X", "Х"},
	{"B", "В"},
	{"M", "М"},
	{"H", "Н"},
}

func latokir(s string) string {
	for _, v := range substitutes {
		s = strings.Replace(s, v[0], v[1], -1)
	}
	return s
}
