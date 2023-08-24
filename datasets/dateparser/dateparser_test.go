package dateparser

import (
	"testing"
	"time"
)

func TestShouldMatchYYYYdashMM(t *testing.T) {
	want := time.Date(2005, time.March, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseDate("2005-03")
	if want != result || err != nil {
		t.Fatalf("Want: %v. Got: %v", want, result)
	}
}

func TestShouldMatchYYYYspaceMMM(t *testing.T) {
	want := time.Date(2005, time.March, 1, 0, 0, 0, 0, time.UTC)
	result, err := ParseDate("2005 MAR")
	if want != result || err != nil {
		t.Fatalf("Want: %v. Got: %v", want, result)
	}
}

func TestShouldMatchDDspaceMMMspaceYY(t *testing.T) {
	want := time.Date(2005, time.March, 4, 0, 0, 0, 0, time.UTC)
	result, err := ParseDate("04 Mar 05")
	if want != result || err != nil {
		t.Fatalf("Want: %v. Got: %v. Err: %v", want, result, err)
	}
}
