package sexpr

import "testing"

type testInterface interface {
	greet() string
}

func TestMarshal(t *testing.T) {
	var i testInterface
	var tests = []struct{
		input interface{}
		wantBytes []byte
		wantErr error
	}{
		{1, []byte("1"), nil},
		{2, []byte("2"), nil},
		{true, []byte("t"), nil},
		{false, nil, nil},
		{1.1, []byte("1.100000"), nil},
		{complex(1, 2), []byte("#C(1.0, 2.0)"), nil},
	}
	for _, test := range tests {
		gotBytes, gotErr := Marshal(test.input)
		if test.wantErr != nil {
			if gotErr == nil {
				t.Errorf("Marshal(%q) should return error.", test.input)
			}
			if gotErr != test.wantErr {
				t.Errorf("Marshal(%q) returned error %q, want %q.", test.input, gotErr, test.wantErr)
			}
		} else if test.wantErr == nil && gotErr != nil {
			t.Errorf("Marshal(%q) should not return error.", test.input)
		}
		if test.wantBytes != nil {
			if gotBytes == nil {
				t.Errorf("Marshal(%q) should return bytes.", test.input)
			}
			if len(gotBytes) != len(test.wantBytes) {
				t.Errorf("Marshal(%q) = %q, want %q.", test.input, gotBytes, test.wantBytes)
			}
			for i := range gotBytes {
				if gotBytes[i] != test.wantBytes[i] {
					t.Errorf("Marshal(%q) = %q, want %q.", test.input, gotBytes, test.wantBytes)
				}
			}
		} else if test.wantBytes == nil && gotBytes != nil {
			t.Errorf("Marshal(%q) should return nil bytes.", test.input)
		}
	}
}
