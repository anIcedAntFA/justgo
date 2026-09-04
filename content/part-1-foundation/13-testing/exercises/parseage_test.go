package exercises

import (
	"errors"
	"testing"
)

func TestParseAge(t *testing.T) {
	t.Skip("Chapter 13 exercise: implement ParseAge, then delete this Skip")

	tests := []struct {
		name    string
		in      string
		want    int
		wantErr error // nil = expect no error; else the sentinel errors.Is must find
	}{
		{"valid", "42", 42, nil},
		{"zero boundary", "0", 0, nil},
		{"upper boundary", "150", 150, nil},
		{"not a number", "abc", 0, ErrNotNumber},
		{"empty", "", 0, ErrNotNumber},
		{"negative", "-1", 0, ErrOutOfRange},
		{"too old", "151", 0, ErrOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAge(tt.in)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseAge(%q) err = %v; want errors.Is %v", tt.in, err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return // on error, the value is defined as 0 but not the point
			}
			if got != tt.want {
				t.Errorf("ParseAge(%q) = %d; want %d", tt.in, got, tt.want)
			}
		})
	}
}
