package charcount

import "testing"

func TestCount(t *testing.T) {
	var tests = []struct {
		input string
		want map[rune]int
	}{
		{"", map[rune]int{}},
		{"a", map[rune]int{'a': 1}},
		{"ab", map[rune]int{'a': 1, 'b':1}},
		{"hi there", map[rune]int{'e': 2, 'h': 2, 'i': 1, 'r': 1, 't': 1, ' ': 1}},
		{"été", map[rune]int{'é': 2, 't': 1}},
	}
	for _, test := range tests {
		got, _, _ := Count(test.input)
		if len(got) != len(test.want) {
			t.Errorf(`Count(%s): len(counts) = %d; want %d`, test.input, len(got), len(test.want))
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf(`Count(%s): counts[%c] = %d; want %d`, test.input, i, got[i], test.want[i])
			}
		}
	}
}

func TestUtfLen(t *testing.T) {
	var tests = []struct {
		input string
		want []int
	}{
		{"a", []int{0, 1, 0, 0, 0}},
		{"hello world", []int{0, 11, 0, 0, 0}},
		{"été", []int{0, 1, 2, 0, 0}},
	}
	for _, test := range tests {
		_, got, _ := Count(test.input)
		for i, val := range got {
			if val != test.want[i] {
				t.Errorf(`Count(%s): utflen[%d] = %d; want %d`, test.input, i, val, test.want[i])
			}
		}
	}
}
