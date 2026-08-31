package exercises

import (
	"slices"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	t.Skip("Chapter 10 exercise: implement WordFrequency, then delete this Skip")

	cases := []struct {
		name string
		in   string
		want []WordCount
	}{
		{"empty", "", []WordCount{}},
		{
			"single word repeated",
			"go go go",
			[]WordCount{{"go", 3}},
		},
		{
			"sorted by count then word",
			"the quick brown fox the lazy dog the end",
			[]WordCount{
				{"the", 3},
				{"brown", 1},
				{"dog", 1},
				{"end", 1},
				{"fox", 1},
				{"lazy", 1},
				{"quick", 1},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := WordFrequency(tc.in); !slices.Equal(got, tc.want) {
				t.Errorf("WordFrequency(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
