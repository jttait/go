package format

import (
	"testing"
	"time"
)

func TestFormatAnyWithInt64(t *testing.T) {
	var tests = []struct {
		input int64
		want string
	}{
		{int64(1), "1"},
	}
	for _, test := range tests {
		got := Any(test.input)
		if got != test.want {
			t.Errorf("formatAtom(%q): got %q, want %q", test.input, got, test.want)
		}
	}
}

func TestFormatAnyWithDuration(t *testing.T) {
	var tests = []struct {
		input time.Duration
		want string
	}{
		{1 * time.Nanosecond, "1"},
		{1 * time.Microsecond, "1000"},
	}
	for _, test := range tests {
		got := Any(test.input)
		if got != test.want {
			t.Errorf("formatAtom(%q): got %q, want %q", test.input, got, test.want)
		}
	}
}
