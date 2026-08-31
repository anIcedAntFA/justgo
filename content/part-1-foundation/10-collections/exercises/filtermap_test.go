package exercises

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	t.Skip("Chapter 10 exercise: implement Filter, then delete this Skip")

	isEven := func(n int) bool { return n%2 == 0 }

	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"some even", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"none match", []int{1, 3, 5}, []int{}},
		{"all match", []int{2, 4}, []int{2, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Filter(tc.in, isEven); !slices.Equal(got, tc.want) {
				t.Errorf("Filter(%v, isEven) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Skip("Chapter 10 exercise: implement Map, then delete this Skip")

	double := func(n int) int { return n * 2 }

	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"doubles", []int{1, 2, 3}, []int{2, 4, 6}},
		{"negatives", []int{-1, 0, 4}, []int{-2, 0, 8}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Map(tc.in, double); !slices.Equal(got, tc.want) {
				t.Errorf("Map(%v, double) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
