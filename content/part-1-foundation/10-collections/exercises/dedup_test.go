package exercises

import (
	"slices"
	"testing"
)

func TestDedup(t *testing.T) {
	t.Skip("Chapter 10 exercise: implement Dedup, then delete this Skip")

	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"no dups", []int{1, 2, 3}, []int{1, 2, 3}},
		{"non-adjacent dups", []int{1, 2, 1, 3, 2}, []int{1, 2, 3}},
		{"all same", []int{5, 5, 5}, []int{5}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Dedup(tc.in); !slices.Equal(got, tc.want) {
				t.Errorf("Dedup(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
