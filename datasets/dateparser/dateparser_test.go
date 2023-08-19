package dateparser

import (
	"testing"
	"time"
)

func TestShouldMatchYYYYdashMM(t *testing.T) {
	want := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseDate("2000-01")
	if want != result || err != nil {
		t.Fatalf("Want: %v. Got: %v", want, result)
	}
}

func TestShouldMatchYYYYspaceMMM(t *testing.T) {
	want := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseDate("2000 JAN")
	if want != result || err != nil {
		t.Fatalf("Want: %v. Got: %v", want, result)
	}
}
