package exercises

import (
	"bytes"
	"testing"
)

func TestWriteReport(t *testing.T) {
	// t.Skip("Chapter 07 exercise: implement WriteReport, then delete this Skip")

	cases := []struct {
		name string
		data []string
		want string
	}{
		{"two items", []string{"alpha", "beta"}, "1. alpha\n2. beta\n"},
		{"one item", []string{"only"}, "1. only\n"},
		{"empty", nil, ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := WriteReport(&buf, tc.data); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := buf.String(); got != tc.want {
				t.Errorf("WriteReport wrote %q, want %q", got, tc.want)
			}
		})
	}
}
