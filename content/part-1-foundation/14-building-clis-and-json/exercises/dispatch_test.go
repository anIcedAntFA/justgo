package exercises

import (
	"errors"
	"testing"
)

func TestDispatch(t *testing.T) {
	t.Skip("Chapter 14 exercise: implement Dispatch, then delete this Skip")

	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"empty defaults to organize", "", "organize", false},
		{"organize", "organize", "organize", false},
		{"stats", "stats", "stats", false},
		{"undo", "undo", "undo", false},
		{"rules", "rules", "rules", false},
		{"unknown", "frobnicate", "", true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Dispatch(tc.in)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Dispatch(%q) error = %v, wantErr %t", tc.in, err, tc.wantErr)
			}
			if tc.wantErr && !errors.Is(err, ErrUnknownCommand) {
				t.Errorf("Dispatch(%q) error = %v, want it to wrap ErrUnknownCommand", tc.in, err)
			}
			if got != tc.want {
				t.Errorf("Dispatch(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
