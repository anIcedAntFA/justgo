package exercises

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	// t.Skip("Chapter 10 exercise: implement Reverse, then delete this Skip")

	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{7}, []int{7}},
		{"odd length", []int{1, 2, 3}, []int{3, 2, 1}},
		{"even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original := slices.Clone(tc.in)

			got := Reverse(tc.in)
			if !slices.Equal(got, tc.want) {
				t.Errorf("Reverse(%v) = %v, want %v", original, got, tc.want)
			}

			// The original must be untouched — Reverse returns a copy.
			if !slices.Equal(tc.in, original) {
				t.Errorf("Reverse mutated its input: got %v, want %v", tc.in, original)
			}
		})
	}
}
