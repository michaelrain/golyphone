package polyphone

import (
	"testing"
)

func TestGolyphone(t *testing.T) {

	a, _ := Encode("Привет")
	b, _ := Encode("Превет")

	if a != b {
		t.Error("should equal")
	}
}
