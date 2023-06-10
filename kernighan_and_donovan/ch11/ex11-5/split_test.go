package split

import (
	"testing"
	"strings"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		s string
		sep string
		want int
	}{
		{"a:b:c", ":", 3},
		{"1,2,3", ",", 3},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d words",
				test.s, test.sep, got, test.want)
		}
	}
}
